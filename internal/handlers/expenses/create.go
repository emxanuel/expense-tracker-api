package expenses_handler

import (
	expenses_dto "expense-tracker/api/dto/expenses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *ExpensesHandler) CreateExpenseHanlder(c *gin.Context) {
	var body expenses_dto.CreateExpenseBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	expense, err := handler.Repo.CreateExpense(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, expense)
}
