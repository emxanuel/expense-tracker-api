package expenses_repository

import (
	expenses_dto "expense-tracker/api/dto/expenses"

	"gorm.io/gorm"
)

type ExpensesRepository struct {
	DB *gorm.DB
}

func NewExpensesRepository(db *gorm.DB) *ExpensesRepository {
	return &ExpensesRepository{DB: db}
}

func (r *ExpensesRepository) GetExpenses() ([]expenses_dto.GetExpensesItem, error) {
	var expenses []expenses_dto.GetExpensesItem

	// if err := r.DB.Find(&expenses).Joins("Category").Error; err != nil {
	// 	return nil, err
	// }

	if err := r.DB.Table("transactions t").
		Joins("inner join categories c on c.id = t.category_id").
		Select("t.amount, t.note, t.date, t.type, c.name as category").
		Scan(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}
