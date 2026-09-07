package routes

import (
	"weather-api/handlers"
	"weather-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, userHandler *handlers.UserHandler) {
	server.POST("/signup", userHandler.Signup)
	server.POST("/login", userHandler.Login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
}
