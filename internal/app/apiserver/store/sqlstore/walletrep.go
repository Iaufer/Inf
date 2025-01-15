package sqlstore

import (
	"Inf/internal/app/model"
	"context"
)

type WalletRepository struct {
	store *Store
}

func (r *WalletRepository) FindByAddress(ctx context.Context, address string) (*model.Wallet, error) {
	return nil, nil
}

func (r *WalletRepository) Update(ctx context.Context, wallet *model.Wallet) error {
	return nil
}
