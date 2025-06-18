package expenses_routes

import "github.com/gin-gonic/gin"

func RegisterExpensesRoutes(router *gin.Engine) {
	expensesRoutes := router.Group("/api/expenses/")
	{
		expensesRoutes.GET("")
	}

}
