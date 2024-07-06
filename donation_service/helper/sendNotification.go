package helper

import (
	"donation_service/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func CreateDonationNotif(body *models.DonationDetailResp) {
	conn, err := amqp.Dial(RABBIT_MQ_ADDR)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		CREATE_DONATION_CH,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(b),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
	log.Printf(" [x] Sent %v", body)
}

func EditDonationNotif(body *models.DonationDetailResp) {
	conn, err := amqp.Dial(RABBIT_MQ_ADDR)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		EDIT_DONATION_CH,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(b),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
	log.Printf(" [x] Sent %v", body)
}
