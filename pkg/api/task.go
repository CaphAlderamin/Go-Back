package api

import (
	. "Rip/pkg/models"
	"Rip/pkg/rabbitmq"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"strconv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func IndexApi(c *gin.Context) {
	body := "It works!"
	err := rabbitmq.AmqpChannel.Publish(
		"",         // exchange
		"GetIndex", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message GetIndex")

	//c.String(http.StatusOK, "GetIndex request received")
}

func GetTasksApi(c *gin.Context) {
	err := rabbitmq.AmqpChannel.Publish(
		"",         // exchange
		"GetTasks", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
		})
	failOnError(err, "Failed to publish a message GetTasks")
}

func AddTaskApi(c *gin.Context) {
	task := Task{}
	err := c.Bind(&task)
	if err != nil {
		log.Fatalln(err)
	}

	marshaledTask, err := json.Marshal(task)
	if err != nil {
		log.Fatalln(err)
	}

	err = rabbitmq.AmqpChannel.Publish(
		"",         // exchange
		"PostTask", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshaledTask,
		})
	failOnError(err, "Failed to publish a message PostTask")

	//c.JSON(http.StatusOK, gin.H{
	//	"data": task,
	//})
}

func ModTaskStatusApi(c *gin.Context) {
	cid := c.Request.FormValue("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	task := Task{Id: id}
	err = c.Bind(&task)
	if err != nil {
		log.Fatalln(err)
	}

	marshaledTask, err := json.Marshal(task)
	if err != nil {
		log.Fatalln(err)
	}

	err = rabbitmq.AmqpChannel.Publish(
		"",              // exchange
		"ModTaskStatus", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshaledTask,
		})
	failOnError(err, "Failed to publish a message ModTaskStatus")

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func ModTaskApi(c *gin.Context) {
	cid := c.Request.FormValue("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	task := Task{Id: id}
	err = c.Bind(&task)
	if err != nil {
		log.Fatalln(err)
	}

	marshaledTask, err := json.Marshal(task)
	if err != nil {
		log.Fatalln(err)
	}

	err = rabbitmq.AmqpChannel.Publish(
		"",        // exchange
		"ModTask", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshaledTask,
		})
	failOnError(err, "Failed to publish a message ModTask")

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func DelTaskApi(c *gin.Context) {
	cid := c.Request.FormValue("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	marshaledId, err := json.Marshal(id)
	if err != nil {
		log.Fatalln(err)
	}

	err = rabbitmq.AmqpChannel.Publish(
		"",           // exchange
		"DeleteTask", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshaledId,
		})
	failOnError(err, "Failed to publish a message DeleteTask")

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}
