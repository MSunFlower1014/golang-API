package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	e := gin.Default()
	e.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong pong"})
	})
	e.Run()
}
