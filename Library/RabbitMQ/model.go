package RabbitMQ

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitConnection *amqp.Connection

type RabbitConfig struct {
	Scheme   string `yaml:"scheme"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Vhost    string `yaml:"vhost"`
}
