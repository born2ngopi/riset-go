package main

import (
	"github.com/streadway/amqp"
)

type RabbitMQ interface{}

type rabbitMq struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	ExchangeType string
	ExchangeName string
}

func NewRabbitMq() (RabbitMQ, error) {

	var url = "amqp://guest:guest@localhost:5672/"

	LogInfo("crating connection with rabbit mq")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	LogInfo("got connection")

	LogInfo("create channel")
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	var exchangeType, exchangeName string = "direct", "learn"

	// exchange
	if err := channel.ExchangeDeclare(
		exchangeName, // exchange name
		exchangeType, // exchange type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, err
	}

	return &rabbitMq{
		Connection:   conn,
		Channel:      channel,
		ExchangeType: exchangeType,
		ExchangeName: exchangeName,
	}, nil
}

func (mq *rabbitMq) Publish(body []byte) error {

	if err := mq.Channel.Publish(
		mq.ExchangeName, // publish to an exchange
		"test-key",      // routing to 0 or more queues
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return err
	}

	return nil
}
