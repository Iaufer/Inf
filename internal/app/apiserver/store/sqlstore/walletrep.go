package sqlstore

import (
	"Inf/internal/app/model"
	"context"
	"fmt"
)

type WalletRepository struct {
	store *Store
}

func (r *WalletRepository) FindByAddress(ctx context.Context, address string) (*model.Wallet, error) {
	var wallet model.Wallet

	if err := r.store.db.WithContext(ctx).First(&wallet, "address = ?", address).Error; err != nil {
		return nil, fmt.Errorf("wallet with address %s was not found: %v", address, err)
	}

	return &wallet, nil
}

func (r *WalletRepository) Update(ctx context.Context, wallet *model.Wallet) error {
	return nil
}
