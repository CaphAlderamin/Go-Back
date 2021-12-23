package transport

import (
	. "Rip/pkg/api"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/tasks", GetTasksApi)

	router.POST("/task", AddTaskApi)

	router.PUT("/task", ModTaskApi)

	router.DELETE("/task", DelTaskApi)

	return router
}
