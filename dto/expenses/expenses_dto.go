package expenses_dto

import "time"

type GetExpensesItem struct {
	Category string
	Amount   float64
	Note     string
	Date     time.Time
	Type     string
}
