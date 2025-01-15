package apiserver

import (
	"Inf/internal/app/apiserver/store/sqlstore"
	"database/sql"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err // сделать ошибку более информативной
	}

	defer db.Close()

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
