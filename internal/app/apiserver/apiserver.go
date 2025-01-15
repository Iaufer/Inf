package apiserver

import (
	"Inf/internal/app/apiserver/store/sqlstore"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err // сделать ошибку более информативной
	}

	defer db.Close()

	if err := migrations(db); err != nil {
		return fmt.Errorf("failed to apply migrations: %v", err)
	} else {
		fmt.Println("Migrations were completed successfully (Note: if they were completed earlier, they will simply be ignored because they were completed earlier)")
	}

	store := sqlstore.New(db)

	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err // inf
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func migrations(db *sql.DB) error {
	files, err := os.ReadDir("../../migrations")

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
		filePath := fmt.Sprintf("../../migrations/%s", file.Name())

		if err := migration(db, filePath); err != nil {
			return fmt.Errorf("failed to apply migration %s: %v", file.Name(), err)
		}
	}

	return nil
}

func migration(db *sql.DB, filePath string) error {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return fmt.Errorf("failed to read migration file %s: %v", filePath, err)
	}

	_, err = db.Exec(string(content))

	if err != nil {
		return fmt.Errorf("failed to apply migration from file %s: %v", filePath, err)
	}

	return nil
}
