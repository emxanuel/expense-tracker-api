package models

import "time"

type Categories struct {
	ID           uint           `gorm:"primaryKey"`
	UserID       uint           `gorm:"column:user_id"`
	Name         string         `gorm:"type:varchar(10)"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamptz"`
	Transactions []Transactions `gorm:"foreignKey:CategoryID"`
	Budgets      []Budgets      `gorm:"foreignKey:CategoryID"`
}

/*
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(100) NOT NULL,
  type VARCHAR(10) CHECK (type IN ('expense', 'income')), -- categorizes the category
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
*/
