package routes

import (
	"weather-api/handlers"
	"weather-api/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.Engine, h *handlers.UserHandler) {
	server.POST("/signup", h.Signup)
	server.POST("/login", h.Login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
}
