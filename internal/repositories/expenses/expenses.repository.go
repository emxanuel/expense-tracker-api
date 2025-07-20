package expenses_repository

import (
	"expense-tracker/api/dto/dtos"
	"expense-tracker/api/models"
	"time"

	"gorm.io/gorm"
)

type ExpensesRepository struct {
	DB *gorm.DB
}

func NewExpensesRepository(db *gorm.DB) *ExpensesRepository {
	return &ExpensesRepository{DB: db}
}

func (r *ExpensesRepository) GetExpenses() ([]dtos.GetExpensesItem, error) {
	var expenses []dtos.GetExpensesItem

	if err := r.DB.Table("transactions t").
		Joins("inner join categories c on c.id = t.category_id").
		Select("t.amount, t.note, t.date, t.type, c.name as category").
		Scan(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpensesRepository) CreateExpense(expense dtos.CreateExpenseBody) (*dtos.GetExpensesItem, error) {
	var createdExpense *dtos.GetExpensesItem
	tx := &models.Transactions{
		Amount:     expense.Amount,
		UserID:     expense.UserID,
		CategoryID: expense.CategoryID,
		Note:       expense.Note,
		Date:       expense.Date,
		Type:       expense.Type,
		CreatedAt:  time.Now(),
	}

	if err := r.DB.Create(&tx).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Table("transactions t").
		Joins("inner join categories c on c.id = t.category_id").
		Select("t.amount, t.note, t.date, t.type, c.name as category").
		Where("t.id = ?", tx.ID).
		Scan(&createdExpense).Error; err != nil {
		return nil, err
	}

	return createdExpense, nil
}
