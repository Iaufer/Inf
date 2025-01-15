package model

type Wallet struct {
	ID      uint   `gorm:"primaryKey"`
	Address string `gorm:"uniqueIndex"`
	Balance float64
}
