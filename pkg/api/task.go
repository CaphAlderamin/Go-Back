package api

import (
	. "Rip/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func GetTasksApi(c *gin.Context) {
	task := Task{}
	tasks, err := task.GetTasks()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func AddTaskApi(c *gin.Context) {
	task := Task{}
	err := c.Bind(&task)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := task.AddTask()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Insert successful! new task id: %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
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

	rv, err := task.ModTask()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Update task %d successful %d", task.Id, rv)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelTaskApi(c *gin.Context) {
	cid := c.Request.FormValue("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	task := Task{Id: id}

	rv, err := task.DelTask()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Delete task %d successful %d", id, rv)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
