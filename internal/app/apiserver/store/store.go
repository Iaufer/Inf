package store

type Store interface {
	Wallet() WalletRepository
	Transaction() TransactionRepository
	// User() UserRepository
}
