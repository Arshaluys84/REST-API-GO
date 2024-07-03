package routes

import (
	"net/http"

	"arsh.com/rest-api/models"
	"arsh.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "please, check your data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Your data has not been saved"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "You have signed up successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "please, check your data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login succesful!", "token": token})
}
