package main

import (
	"Rip/pkg/models"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//conn, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	_, err = amqpChannel.QueueDeclare(
		"GetIndex",
		false,
		false,
		false,
		false,
		nil)
	handleError(err, "Could not declare `GetIndex` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	task := models.Task{TaskName: "Task", TaskDescription: "Ебать копать"}

	body, err := json.Marshal(task)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	//for i := 0; i < 1; i++ {
	err = amqpChannel.Publish(
		"",
		"GetIndex",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
	//}

	log.Printf("GetIndex send")
}
