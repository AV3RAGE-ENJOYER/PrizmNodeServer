package handlers

import (
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
}
