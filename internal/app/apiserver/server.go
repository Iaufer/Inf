package apiserver

import (
	"Inf/internal/app/apiserver/store"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/api/send", s.handleSend()).Methods("POST")

}

func (s *server) handleSend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			From   string `json:"from"`
			To     string `json:"to"`
			Amount string `json:"amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// if err := s.store

	}
}
