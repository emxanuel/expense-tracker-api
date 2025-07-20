package expenses_handler

import expenses_repository "expense-tracker/api/internal/repositories/expenses"

type ExpensesHandler struct {
	Repo *expenses_repository.ExpensesRepository
}

func NewExpensesHandler(repo *expenses_repository.ExpensesRepository) *ExpensesHandler {
	return &ExpensesHandler{Repo: repo}
}
