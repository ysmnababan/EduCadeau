package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"registry_service/models"

	"github.com/streadway/amqp"
)

func CreateRegistryNotif(body *models.Registry) {
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
		CREATE_REGISTRY_CH,
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
