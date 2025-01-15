package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" //...
)

type Store struct {
	db *sql.DB
	// userRepository  *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
