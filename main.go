package main

import (
	"Go-RestApi/db"
	"Go-RestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Fetching Data Fail"})
		return
	}
	// HTTP 요청에 대한 응답을 직접 처리하는 역할하기때문에 리턴 안해도됨
	context.JSON(http.StatusOK, events) //json  형태로 보내겟다 리턴을 안해도됨
}

func createEvents(context *gin.Context) {
	var event models.Event
	// HTTP 요청 바디의 JSON 데이터를 event 구조체에 바인딩합니다.
	// &event는 event 구조체의 메모리 주소를 의미하며,
	// ShouldBindJSON 함수는 이 주소를 통해 JSON 데이터를 구조체에 직접 채워 넣습니다.
	// 즉, 클라이언트에서 전달된 JSON 데이터의 각 필드가 event 구조체의 해당 필드에 매핑됩니다.
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing Error data"})
		return
	}
	//하기의 코드는 추후 DB 연결 후 바꿀 예정
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Save Data Fail"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Success Create", "event": event})
}
