package routes

import (
	"Go-RestApi/models"
	"Go-RestApi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request Data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func logIn(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request User Data"})
		return
	}
	err = user.ValidateCredential() // 유저객체에 유효성 검사에서 쓰인 유저의 이메일과 패스워드와 아이디와 같은 전부 필요한 값들을 가지고 있음

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credential"})
	}

	// 위에서 user 객체에 필요한 정보를 다 가지고 있기때문에 파라미터로 받아서 쓸수가 있음
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not generate token"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User Login Success", "token": token})
}
