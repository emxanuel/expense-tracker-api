package routes

import (
	budgets_routes "expense-tracker/api/api/routes/budgets"
	expenses_routes "expense-tracker/api/api/routes/expenses"
	default_handler "expense-tracker/api/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	expenses_routes.RegisterExpensesRoutes(router, db)
	budgets_routes.RegisterBudgetsRoutes(router, db)
	{
		router.GET("/", default_handler.Greetings)
	}
}
