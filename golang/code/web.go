package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/message/:content", func(c *gin.Context) {
		content := c.Param("content")
		c.String(http.StatusOK, `{"message": "%s"}`, content)
	})

	r.Run(":8000")
}
