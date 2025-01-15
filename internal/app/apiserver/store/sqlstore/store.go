package sqlstore

import (
	_ "github.com/lib/pq" //...
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
	// userRepository  *UserRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}
