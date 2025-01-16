package apiserver

import (
	"Inf/internal/app/apiserver/store"
	"encoding/json"
	"fmt"
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
			Amount int    `json:"amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		fmt.Println(req)

		ctx := r.Context()

		// err := s.store.Transaction()

		fromWallet, err := s.store.Wallet().FindByAddress(ctx, req.From)

		if err != nil {
			return
			// return fmt.Errorf("sender wallet not found")
		}

		toWallet, err := s.store.Wallet().FindByAddress(ctx, req.To)

		if err != nil {
			// return fmt.Errorf("sender wallet not found")
			return
		}

		if fromWallet.Balance < float64(req.Amount) {
			fmt.Println("insufficient funds")
			return
		}

		fromWallet.Balance -= float64(req.Amount)
		toWallet.Balance += float64(req.Amount)

		if err := s.store.Wallet().Update(ctx, fromWallet); err != nil {
			fmt.Println("Ошибка в отправителе кошелька")
			return
		}

		if err := s.store.Wallet().Update(ctx, toWallet); err != nil {
			fmt.Println("Ошибка в получателе кошелька")
			return
		}

		fmt.Println("fromWallet: ", fromWallet.Balance)
		fmt.Println("toWallet: ", toWallet.Balance)

		// fmt.Println("fromWallet: ", fromWallet.Address, "переслать: ", req.Amount, "toWallet: ", toWallet.Address)

	}
}
