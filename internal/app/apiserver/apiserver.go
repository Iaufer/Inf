package apiserver

import (
	"Inf/internal/app/apiserver/store/sqlstore"
	"fmt"
	"net/http"
	"os"
	"sort"

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

	fmt.Println(files)

	if err != nil {
		return fmt.Errorf("failed to read migrations folder: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() { //skip if not is files
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
