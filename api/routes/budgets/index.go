package budgets_routes

import (
	budgets_handler "expense-tracker/api/internal/handlers/budgets"
	budgets_repository "expense-tracker/api/internal/repositories/budgets"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBudgetsRoutes(router *gin.Engine, db *gorm.DB) {
	budgetsRoutes := router.Group("/api/budgets/")
	budgetsRepository := budgets_repository.NewBudgetsRepository(db)
	budgetsHandler := budgets_handler.NewBudgetsHandler(budgetsRepository)
	{
		budgetsRoutes.GET("", budgetsHandler.GetBudgetsHanlder)
	}
}
