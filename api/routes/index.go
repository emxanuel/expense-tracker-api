package routes

import (
	expenses_routes "expense-tracker/api/api/routes/expenses"
	default_handler "expense-tracker/api/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	expenses_routes.RegisterExpensesRoutes(router, db)
	{
		router.GET("/", default_handler.Greetings)
	}
}
