package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var AmqpChannel *amqp.Channel
var Connection *amqp.Connection

func InitRabbitMQ() {
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	Connection = conn

	ch, err := Connection.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	AmqpChannel = ch
	err = AmqpChannel.Qos(1, 0, false)
	failOnError(err, "Could not configure QoS")

	_, err = ch.QueueDeclare(
		"GetIndex", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue GetIndex")

	_, err = ch.QueueDeclare(
		"GetTasks", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue GetTasks")

	_, err = ch.QueueDeclare(
		"PostTask", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue PostTask")

	_, err = ch.QueueDeclare(
		"DeleteTask", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue DeleteTask")

	_, err = ch.QueueDeclare(
		"ModTask", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue ModTask")

	_, err = ch.QueueDeclare(
		"ModTaskStatus", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue ModTaskStatus")

}
