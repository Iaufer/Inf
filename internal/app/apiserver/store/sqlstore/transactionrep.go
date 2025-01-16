package sqlstore

import (
	"Inf/internal/app/model"
	"context"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	store *Store
}

func (r *TransactionRepository) Create(ctx context.Context, tr *model.Transaction) error {
	return nil
}

func (r *TransactionRepository) DB() *gorm.DB {
	return r.store.db
}
