package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/book_usage_statistic_items", func(c *gin.Context){
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}



