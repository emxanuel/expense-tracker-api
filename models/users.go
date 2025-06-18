package models

import "time"

type Users struct {
	ID           uint           `gorm:"primaryKey"`
	FirstName    string         `gorm:"column:first_name;type:varchar(100)"`
	LastName     string         `gorm:"column:last_name;type:varchar(100)"`
	Email        string         `gorm:"type:text"`
	PasswordHash string         `gorm:"column:password_hash;type:text"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamptz"`
	Transactions []Transactions `gorm:"foreignKey:UserID"`
	Budgets      []Budgets      `gorm:"foreignKey:UserID"`
	Categories   []Categories   `gorm:"foreignKey:UserID"`
}

/*
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(100),
	last_name VARCHAR(100),
  email VARCHAR(100) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
*/
