package Publisher

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"mobile-banking-service/Constants"
	"mobile-banking-service/Controllers/ControllersPublisher"
	"mobile-banking-service/Library/Helper"
	"mobile-banking-service/Library/RabbitMQ"
)

func (r *rabbitPublisher) published(channel *amqp.Channel) (err error) {
	err = channel.Publish(
		r.Publisher.exchange,
		r.Publisher.key,
		r.Publisher.mandatory,
		r.Publisher.immediate,
		r.Publisher.publishing,
	)
	return
}

func InitRabbitPublish(publisher ControllersPublisher.DataPublish) Publisher {
	return &rabbitPublisher{
		Publisher: &DataPublisher{
			exchange: string(publisher.ExchangeName),
			key:      string(publisher.Key),
			publishing: amqp.Publishing{
				ContentType: Constants.ContentTypeJSON,
				Body:        publisher.DataRequest,
			},
		},
	}
}

func (r *rabbitPublisher) newChannel() (channel *amqp.Channel, err error) {
	channel, err = RabbitMQ.RabbitConnection.Channel()
	return
}

func (r *rabbitPublisher) PublishNonRPC() (err error) {
	channel, err := r.newChannel()
	if err != nil {
		return
	}

	defer channel.Close()
	err = r.published(channel)
	return
}

func (r *rabbitPublisher) PublishRPC() (data Helper.ResponsePublisher) {
	channel, err := r.newChannel()
	if err != nil {
		data.Status = err.Error()
		return
	}
	corrId := uuid.NewV4().String()
	q, err := channel.QueueDeclare(
		r.queueTemp(),
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		data.Status = err.Error()
		return
	}
	defer channel.Close()
	msgs, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		data.Status = err.Error()
		return
	}
	r.Publisher.publishing.CorrelationId = corrId
	r.Publisher.publishing.ReplyTo = q.Name
	err = r.published(channel)
	if err != nil {
		data.Status = err.Error()
		return
	}
	for d := range msgs {
		if corrId == d.CorrelationId {
			json.Unmarshal(d.Body, &data)
			return
		}
	}
	return
}

func (r *rabbitPublisher) queueTemp() string {
	uuid := uuid.NewV4().String()
	return fmt.Sprintf("%s-%s", r.Publisher.key, uuid)
}
