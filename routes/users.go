package routes

import (
	"Go-RestApi/models"
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
	err = user.ValidateCredential()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credential"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Login Success"})
}
