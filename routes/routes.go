package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)         //다건 조회
	server.GET("/event/:id", getEvent)       //단건 조회
	server.POST("/event", createEvents)      //데이터 생성
	server.PUT("/event/:id", updateEvent)    //데이터 수정
	server.DELETE("/event/:id", deleteEvent) //데이터 삭제

	server.POST("/signUp", signUp) //유저 회원가입
	server.POST("/login", logIn)   //유저 로그인
}
