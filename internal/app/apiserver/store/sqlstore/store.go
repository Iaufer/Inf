package sqlstore

import (
	"Inf/internal/app/apiserver/store"

	_ "github.com/lib/pq" //...
	"gorm.io/gorm"
)

type Store struct {
	db                    *gorm.DB
	walletRepository      *WalletRepository
	transactionRepository *TransactionRepository
	// userRepository  *UserRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Wallet() store.WalletRepository {
	if s.walletRepository != nil {
		return s.walletRepository
	}

	s.walletRepository = &WalletRepository{
		store: s,
	}

	return s.walletRepository
}

func (s *Store) Transaction() store.TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store: s,
	}

	return s.transactionRepository
}
