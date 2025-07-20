package budgets_handler

import budgets_repository "expense-tracker/api/internal/repositories/budgets"

type BudgetsHandler struct {
	Repo *budgets_repository.BudgetsRepository
}

func NewBudgetsHandler(repo *budgets_repository.BudgetsRepository) *BudgetsHandler {
	return &BudgetsHandler{Repo: repo}
}
