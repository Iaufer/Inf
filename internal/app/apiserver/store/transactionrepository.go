package store

import (
	"Inf/internal/app/model"
	"context"
)

type TransactionRepository interface {
	Create(ctx context.Context, tr *model.Transaction) error
}
