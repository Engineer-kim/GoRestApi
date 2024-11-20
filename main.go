package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	// HTTP 요청에 대한 응답을 직접 처리하는 역할하기때문에 리턴 안해도됨
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) //json  형태로 보내겟다 리턴을 안해도됨
}
