package budgets_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *BudgetsHandler) GetBudgetsHanlder(c *gin.Context) {
	budgets, err := handler.Repo.GetBudgets()

	if err != nil {
		log.Fatal("Error getting budgets: ", err)
	}

	c.JSON(http.StatusOK, budgets)
}
