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
	return r.store.db.WithContext(ctx).Create(tr).Error
}

func (r *TransactionRepository) GetTr(count int, trs *[]*model.Transaction) error {
	return r.store.db.Order("created_at DESC").Limit(count).Find(trs).Error
}

func (r *TransactionRepository) DB() *gorm.DB {
	return r.store.db
}
