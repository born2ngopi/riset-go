package main

import (
	"github.com/streadway/amqp"
)

func (app *APP) SendNotif() {

	var done = make(chan bool)

	LogInfo("declare queue")
	queue, err := app.RabbitMQ.QueueDeclare()
	if err != nil {
		LogError(err, map[string]interface{}{
			"Message": "cannot declare queue",
		})
		return
	}

	LogInfo("bind queue")
	if err := app.RabbitMQ.QueueBind(queue); err != nil {
		LogError(err, map[string]interface{}{
			"Message": "cannot bind queue",
		})
		return
	}

	LogInfo("waiting consume data ...")
	deliveries, err := app.RabbitMQ.Consume(queue)
	if err != nil {
		LogError(err, map[string]interface{}{
			"Message": "cannot consume data",
		})
		return
	}

	go handler(deliveries, done)

	// break
	if <-done {
		return
	}
}

func handler(deliveries <-chan amqp.Delivery, done chan bool) {

	for delivery := range deliveries {
		LogData(map[string]interface{}{
			"len body":     len(delivery.Body),
			"delivery tag": delivery.ConsumerTag,
			"data":         string(delivery.Body),
		})

		delivery.Ack(false)
	}
}
