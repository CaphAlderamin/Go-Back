package transport

import (
	. "Rip/pkg/api"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLFiles("index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws", WebSocketMessageReceiver)

	return router
}
