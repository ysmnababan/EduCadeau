package handler

import (
	"encoding/json"
	"log"
	"notification_service/helper"
	"notification_service/models"
	"time"

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
			SendToMail(data, "REGISTER NEW USER")
			time.Sleep(5 * time.Second)

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

			SendToMail(data, "EDIT USER DATA")
			time.Sleep(5 * time.Second)

		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C\n", helper.USER_EDIT_CHANNEL)
	<-forever
}

func NotifyCreateDonation() {
	conn, err := amqp.Dial(helper.RABBIT_MQ_ADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		helper.CREATE_DONATION_CH,
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
			data := models.DonationDetailResp{}

			err := json.Unmarshal(d.Body, &data)

			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Received a message from %s: %v", helper.CREATE_DONATION_CH, data)
			// Process the message here
			SendToMail(data, "CREATE DONATION DATA")
			time.Sleep(5 * time.Second)

		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C\n", helper.CREATE_DONATION_CH)
	<-forever
}

func NotifyEditDonation() {
	conn, err := amqp.Dial(helper.RABBIT_MQ_ADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		helper.EDIT_DONATION_CH,
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
		log.Println("here")
		for d := range msgs {
			data := models.DonationDetailResp{}

			err := json.Unmarshal(d.Body, &data)

			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Received a message from %s: %v", helper.EDIT_DONATION_CH, data)
			// Process the message here
			SendToMail(data, "EDIT DONATION DATA")
			time.Sleep(5 * time.Second)
		}
	}()

	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C\n", helper.EDIT_DONATION_CH)
	<-forever
}
