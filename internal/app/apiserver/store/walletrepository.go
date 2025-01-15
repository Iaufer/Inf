package store

import (
	"Inf/internal/app/model"
	"context"
)

type WalletRepository interface {
	FindByAddress(ctx context.Context, address string) (*model.Wallet, error)
	Update(ctx context.Context, wallet *model.Wallet) error
}
