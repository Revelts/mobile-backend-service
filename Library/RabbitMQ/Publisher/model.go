package Publisher

import (
	"mobile-banking-service/Library/Helper"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	published(channel *amqp.Channel) error
	PublishNonRPC() (err error)
	PublishRPC() (data Helper.ResponsePublisher)
	newChannel() (channel *amqp.Channel, err error)
}

type DataPublisher struct {
	exchange   string
	key        string
	mandatory  bool
	immediate  bool
	publishing amqp.Publishing
}

type rabbitPublisher struct {
	Publisher *DataPublisher
}
