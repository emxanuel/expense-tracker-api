package budgets_repository

import (
	"expense-tracker/api/dto/dtos"

	"gorm.io/gorm"
)

type BudgetsRepository struct {
	DB *gorm.DB
}

func NewBudgetsRepository(db *gorm.DB) *BudgetsRepository {
	return &BudgetsRepository{DB: db}
}

func (r *BudgetsRepository) GetBudgets() ([]dtos.GetBudgetsItem, error) {
	var budgets []dtos.GetBudgetsItem

	if err := r.DB.Table("budgets b").
		Select("TRIM(to_char(b.month, 'Month')) as month, b.amount").
		Scan(&budgets).Error; err != nil {
		return nil, err
	}

	return budgets, nil
}
