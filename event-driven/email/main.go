package main

import (
	"log"
	"runtime"
)

type APP struct {
	RabbitMQ *rabbitMq
}

func main() {

	app := &APP{}

	mq, err := NewRabbitMq()
	if err != nil {
		LogError(err, map[string]interface{}{
			"Message":       "cannot create intance rabbit mq",
			"how to solve?": "maybe you need to check whether rabbitmq is running or not",
		})
		return
	}

	// register recovery function
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("panic running process: %v\n%s\n", r, buf)
		}
	}()

	app.RabbitMQ = mq.(*rabbitMq)

	app.SendNotif()

	app.RabbitMQ.Connection.Close()
	app.RabbitMQ.Channel.Close()

	LogInfo("done")
}
