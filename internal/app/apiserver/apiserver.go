package apiserver

import (
	"Inf/internal/app/apiserver/store/sqlstore"
	"Inf/internal/app/model"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"sort"

	"crypto/rand"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}

	if err := migrations(db); err != nil {
		return fmt.Errorf("failed to apply migrations: %v", err)
	} else {
		fmt.Println("Migrations applied successfully!")
	}

	if err := initializeWallets(db); err != nil {
		return fmt.Errorf("failed to init wallets: %v", err)
	}

	store := sqlstore.New(db)

	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURl), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func migrations(db *gorm.DB) error {
	files, err := os.ReadDir("migrations")

	// fmt.Println(files)

	if err != nil {
		return fmt.Errorf("failed to read migrations folder: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := fmt.Sprintf("migrations/%s", file.Name())

		cont, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", filePath, err)
		}

		if err := db.Exec(string(cont)).Error; err != nil {
			return fmt.Errorf("failed to apply migration %s: %v", file.Name(), err)
		}
	}

	return nil
}

func initializeWallets(db *gorm.DB) error {
	var count int64

	if err := db.Model(&model.Wallet{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check wallets in db: %v", err)
	}

	if count != 0 {
		return nil
	}

	for i := 0; i < 10; i++ {
		addr, err := generateRandAddr()

		if err != nil {
			return fmt.Errorf("faield to generate addr: %v", err)
		}

		wallet := &model.Wallet{
			Address: addr,
			Balance: 100.0,
		}

		fmt.Println(wallet)

		if err := db.Create(wallet).Error; err != nil {
			return fmt.Errorf("failed to create wallet: %v", err)
		}
	}

	return nil
}

func generateRandAddr() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	addr := hex.EncodeToString(bytes)

	return addr, nil
}
