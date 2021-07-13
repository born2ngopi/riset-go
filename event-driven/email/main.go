package main

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
	app.RabbitMQ = mq.(*rabbitMq)

	app.SendNotif()

	app.RabbitMQ.Connection.Close()
	app.RabbitMQ.Channel.Close()

	LogInfo("done")
}
