package middlewares

import (
	"net/http"

	"arsh.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	user, error := utils.VerifyToken(token)

	if error != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized Token"})
		return
	}
	context.Set("userId", user)
	context.Next()
}
