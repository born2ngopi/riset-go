package main

import (
	"log"
	"net/http"
	"runtime"
)

const (
	PORT    = ":7777"
	VERSION = "version 1.0.0"
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
	}
	app.RabbitMQ = mq.(*rabbitMq)

	http.HandleFunc("/", app.Version)
	http.HandleFunc("/message", app.Message)

	LogInfo("running on port :7777")

	// register recovery function
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("panic running process: %v\n%s\n", r, buf)
		}
	}()

	http.ListenAndServe(PORT, nil)
}
