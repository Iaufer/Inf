package model

import "time"

type Transaction struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	From      string    `gorm:"not null"`
	To        string    `gorm:"not null"`
	Amount    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
