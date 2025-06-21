package expenses_dto

import "time"

type GetExpensesItem struct {
	Category string
	Amount   float64
	Note     string
	Date     time.Time
	Type     string
}

type CreateExpenseBody struct {
	CategoryID uint
	UserID     uint
	Amount     float64
	Note       string
	Date       time.Time
	Type       string
}
