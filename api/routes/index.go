package routes

import (
	default_handler "expense-tracker/api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	{
		router.GET("/", default_handler.Greetings)
	}
}
