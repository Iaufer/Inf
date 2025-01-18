package store

import (
	"Inf/internal/app/model"
	"context"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tr *model.Transaction) error
	GetTr(count int, trs *[]*model.Transaction) error
	DB() *gorm.DB
}
