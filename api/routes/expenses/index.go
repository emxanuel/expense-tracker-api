package expenses_routes

import (
	expenses_handler "expense-tracker/api/internal/handlers/expenses"
	expenses_repository "expense-tracker/api/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterExpensesRoutes(router *gin.Engine, db *gorm.DB) {
	expensesRoutes := router.Group("/api/expenses/")
	expensesRepository := expenses_repository.NewExpensesRepository(db)
	expensesHandler := expenses_handler.NewExpensesHandler(expensesRepository)
	{
		expensesRoutes.GET("", expensesHandler.GetExpensesHandler)
		expensesRoutes.POST("", expensesHandler.CreateExpenseHanlder)
	}

}
