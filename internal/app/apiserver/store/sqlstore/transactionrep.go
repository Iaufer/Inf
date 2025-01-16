package sqlstore

import (
	"Inf/internal/app/model"
	"context"
)

type TransactionRepository struct {
	store *Store
}

func (r *TransactionRepository) Create(ctx context.Context, tr *model.Transaction) error {

	return nil
}
