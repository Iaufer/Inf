package model

import "time"

type Transaction struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	From      string    `gorm:"column:from_address;not null"`
	To        string    `gorm:"column:to_address;not null"`
	Amount    float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
