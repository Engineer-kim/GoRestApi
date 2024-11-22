package routes

import (
	"Go-RestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**각 엔드포인트 별 세부로직*/

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

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not parse event id"})
		return
	}

	context.JSON(http.StatusOK, event)
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

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could Not parse event id"})
		return
	}

	_, err = models.GetEventByID(eventId) //수정할 이벤트가 DB에 있는지 체크 하는 부분

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Fetch eventData"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) //클라이언트에서 받아온 정보를 UpdateEvent 에 할당
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing Error data(During Update)"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Update Data Fail"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success Update"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could Not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId) //삭제할 이벤트가 DB에 있는지 체크 하는 부분

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Fetch eventData"})
		return
	}
	// 삭제 할 이벤트아이디(식별자)를 안보내도되는 이유는
	//event, err := models.GetEventByID(eventId) 이미 여기서 이벤트 아이디를 비롯한 정보를 담아서 보내고 있기 때문
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Delete Data Fail"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success Delete"})

}
