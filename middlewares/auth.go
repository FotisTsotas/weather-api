package middlewares

import (
	"net/http"

	"weather-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token required"})
		return
	}

	user, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization token"})
		return
	}
	context.Set("user", user)
	context.Next()
}

func GetUser(context *gin.Context) (jwt.MapClaims, bool) {
	user, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return nil, false
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return nil, false
	}

	return claims, true

}

func GetUserId(context *gin.Context) *int {
	user, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return nil
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return nil
	}

	userId := int(claims["userId"].(float64))
	return &userId

}
