package models

import "time"

type Budgets struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"column:user_id"`
	CategoryID uint      `gorm:"column:user_id"`
	Month      time.Time `gorm:"type:date"`
	Amount     float64   `gorm:"type:decimal(10,2)"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamptz"`
}

/*
CREATE TABLE budgets (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE,
  month DATE NOT NULL, -- first day of month, e.g., 2025-06-01
  amount DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(user_id, category_id, month)
);
*/
