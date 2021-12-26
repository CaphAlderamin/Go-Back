package api

import (
	. "Rip/pkg/models"
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
		"data": tasks,
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
	task.Id = int(ra)

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
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

	task, err = task.ModTaskStatus()
	if err != nil {
		log.Fatalln(err)
	}

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

	task, err = task.ModTask()
	if err != nil {
		log.Fatalln(err)
	}

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

	task := Task{Id: id}

	_, err = task.DelTask()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}
