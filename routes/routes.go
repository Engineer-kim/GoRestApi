package routes

import (
	"Go-RestApi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)   //다건 조회
	server.GET("/event/:id", getEvent) //단건 조회
	//JWT 토큰 있어야 접근 가능한 URL (세션이 있어야함)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", createEvents)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

	//server.POST("/event", createEvents)      //데이터 생성
	//server.PUT("/event/:id", updateEvent)    //데이터 수정
	//server.DELETE("/event/:id", deleteEvent) //데이터 삭제

	server.POST("/signUp", signUp) //유저 회원가입
	server.POST("/login", logIn)   //유저 로그인
}
