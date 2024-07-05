package handler

import (
	"encoding/json"
	"log"
	"notification_service/helper"
	"notification_service/models"

	"github.com/streadway/amqp"
)

func NotifyUserRegister() {
	conn, err := amqp.Dial(helper.RABBIT_MQ_ADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		helper.USER_REGISTER_CHANNEL,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			data := models.UserResponse{}

			err := json.Unmarshal(d.Body, &data)

			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Received a message from %s: %v", helper.USER_REGISTER_CHANNEL, data)
			// Process the message here
			SendToMail(data)
		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C\n", helper.USER_REGISTER_CHANNEL)
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NotifyUserEditData() {
	conn, err := amqp.Dial(helper.RABBIT_MQ_ADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		helper.USER_EDIT_CHANNEL,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			data := models.UserDetailResponse{}

			err := json.Unmarshal(d.Body, &data)

			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Received a message from %s: %v", helper.USER_EDIT_CHANNEL, data)
			// Process the message here
			SendToMail(data)
		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C\n", helper.USER_EDIT_CHANNEL)
	<-forever
}
