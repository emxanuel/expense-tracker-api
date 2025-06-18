package models

import "time"

type Transactions struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"column:user_id"`
	CategoryID uint      `gorm:"column:category_id"`
	Amount     float64   `gorm:"type:decimal(10,2)"`
	Note       string    `gorm:"type:text"`
	Date       time.Time `gorm:"type:timestamptz"`
	Type       string    `gorm:"type:varchar(10);check:type in ('expense', 'income')"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamptz"`
}

/*
CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
  amount DECIMAL(10, 2) NOT NULL,
  note TEXT,
  date DATE NOT NULL,
  type VARCHAR(10) CHECK (type IN ('expense', 'income')) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
*/
