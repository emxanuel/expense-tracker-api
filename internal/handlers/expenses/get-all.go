package expenses_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *ExpensesHandler) GetExpensesHandler(c *gin.Context) {
	expenses, err := handler.Repo.GetExpenses()

	if err != nil {
		log.Fatal("Error getting expenses: ", err)
	}

	c.JSON(http.StatusOK, expenses)
}
