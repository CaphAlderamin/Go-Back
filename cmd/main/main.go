package main

import (
	db "Rip/pkg/database"
	"Rip/pkg/models"
	"Rip/pkg/rabbitmq"
	. "Rip/pkg/transport"
)

func main() {
	rabbitmq.InitRabbitMQ()
	defer rabbitmq.Connection.Close()
	defer rabbitmq.AmqpChannel.Close()

	db.InitDB()
	defer db.MySqlDB.Close()

	models.GetIndex()
	models.GetTasks()

	models.DelTask()
	models.AddTask()
	models.ModTaskStatus()
	models.ModTask()

	router := InitRoute()
	router.Run(":8080")
}
