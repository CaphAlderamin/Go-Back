package models

import (
	db "Rip/pkg/database"
	"Rip/pkg/rabbitmq"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

type Task struct {
	Id              int    `json:"id" form:"id"`
	TaskName        string `json:"taskName" form:"taskName"`
	TaskDescription string `json:"taskDescription" form:"taskDescription"`
	TaskStatus      bool   `json:"taskStatus" form:"taskStatus"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Get index string
func GetIndex() {
	msgs, err := rabbitmq.AmqpChannel.Consume(
		"GetIndex",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer GetIndex")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}

//Get all tasks list
func GetTasks() {
	msgs, err := rabbitmq.AmqpChannel.Consume(
		"GetTasks",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer GetTasks")

	go func() {
		for _ = range msgs {
			tasks := make([]Task, 0)
			rows, err := db.MySqlDB.Query("SELECT id, taskName, taskDescription, taskStatus FROM checkList")
			defer rows.Close()
			if err != nil {
				return
			}

			for rows.Next() {
				var task Task
				err := rows.Scan(&task.Id, &task.TaskName, &task.TaskDescription, &task.TaskStatus)
				failOnError(err, "Rows scan from bd is failed")
				tasks = append(tasks, task)
			}
			failOnError(err, "Rows next from bd is failed")

			marshaledTasks, err := json.Marshal(tasks)
			if err != nil {
				log.Fatalln(err)
			}

			err = rabbitmq.AmqpChannel.Publish(
				"",             // exchange
				"GetTasksResp", // routing key
				false,          // mandatory
				false,          // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        marshaledTasks,
				})
			failOnError(err, "Failed to publish a message GetTasksResp")
		}
	}()
}

//Create new task (Body raw JSON)
func AddTask() {
	msgs, err := rabbitmq.AmqpChannel.Consume(
		"PostTask",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer GetIndex")

	go func() {
		for d := range msgs {
			task := Task{}
			err := json.Unmarshal(d.Body, &task)
			failOnError(err, "Failed to unmarshal task")

			_, err = db.MySqlDB.Exec(
				"INSERT INTO checkList(taskName, taskDescription, taskStatus) VALUES (?, ?, ?)",
				task.TaskName, task.TaskDescription, task.TaskStatus)
			failOnError(err, "Failed to write task in database")
		}
	}()
}

//Update the task status by id (Query Params + Body raw JSON)
func ModTaskStatus() {
	msgs, err := rabbitmq.AmqpChannel.Consume(
		"ModTaskStatus",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer ModTaskStatus")

	go func() {
		for d := range msgs {
			task := Task{}
			err := json.Unmarshal(d.Body, &task)

			_, err = db.MySqlDB.Exec(
				"UPDATE checkList SET taskStatus=? WHERE id=?",
				task.TaskStatus, task.Id)
			failOnError(err, "Failed to rewrite status task in database")
		}
	}()
}

//Update the task by id (Query Params + Body raw JSON)
func ModTask() {
	msgs, err := rabbitmq.AmqpChannel.Consume(
		"ModTask",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer ModTask")

	go func() {
		for d := range msgs {
			task := Task{}
			err := json.Unmarshal(d.Body, &task)

			_, err = db.MySqlDB.Exec(
				"UPDATE checkList SET taskName=?, taskDescription=? WHERE id=?",
				task.TaskName, task.TaskDescription, task.Id)
			failOnError(err, "Failed to rewrite task in database")
		}
	}()
}

//Delete the task by id
func DelTask() {
	msgs, err := rabbitmq.AmqpChannel.Consume(
		"DeleteTask",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer DeleteTask")

	go func() {
		for d := range msgs {
			id, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to decode DeleteTask body")

			_, err = db.MySqlDB.Exec("DELETE FROM checkList WHERE id=?", id)
			failOnError(err, "Failed to delete task in database")
		}
	}()
}

func getResTask(id int) (resTask Task, err error) {
	rows, err := db.MySqlDB.Query("SELECT id, taskName, taskDescription, taskStatus FROM checkList WHERE id=?", id)
	defer rows.Close()
	if err != nil {
		return
	}

	rows.Next()
	err = rows.Scan(&resTask.Id, &resTask.TaskName, &resTask.TaskDescription, &resTask.TaskStatus)
	if err != nil {
		return
	}

	return
}
