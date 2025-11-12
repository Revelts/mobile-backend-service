package RabbitMQ

import (
	"bytes"
	"log"

	"github.com/getsentry/sentry-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (config RabbitConfig) InitConnection(ApplicationName string) {
	var err error
	var buffer bytes.Buffer

	buffer.WriteString(config.Scheme + "://")
	buffer.WriteString(config.Username + ":" + config.Password)
	buffer.WriteString("@")
	buffer.WriteString(config.Host + ":" + config.Port)
	buffer.WriteString(config.Vhost)
	connectionString := buffer.String()

	rabbitConfig := amqp.Config{
		Properties: amqp.Table{
			"connection_name": ApplicationName,
		},
	}
	RabbitConnection, err = amqp.DialConfig(connectionString, rabbitConfig)
	if err != nil {
		sentry.CaptureException(err)
		panic(err.Error())
		return
	}

	log.Println("ENGINE rabbit start....")
}
