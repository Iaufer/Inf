package apiserver

import (
	"Inf/internal/app/apiserver/store"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
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
			From   string  `json:"from" validate:"required"`
			To     string  `json:"to" validate:"required`
			Amount float32 `json:"amount" validate:"required,gt=0"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		validate := validator.New()

		if err := validate.Struct(req); err != nil {
			http.Error(w, fmt.Sprintf("validate failed: %v", err), http.StatusBadRequest)
			return
		}

		fmt.Println(req)

		ctx := r.Context()

		err := s.store.Transaction().DB().Transaction(func(tx *gorm.DB) error {
			fromWallet, err := s.store.Wallet().FindByAddress(ctx, req.From)

			if err != nil {
				return err
				// return fmt.Errorf("sender wallet not found")
			}

			toWallet, err := s.store.Wallet().FindByAddress(ctx, req.To)

			if err != nil {
				// return fmt.Errorf("sender wallet not found")
				return err
			}

			if fromWallet.Balance < float64(req.Amount) {
				http.Error(w, "insufficient funds", http.StatusBadRequest)
				return fmt.Errorf("insufficient funds")
			}

			fromWallet.Balance -= float64(req.Amount)
			toWallet.Balance += float64(req.Amount)

			if err := s.store.Wallet().Update(ctx, fromWallet); err != nil {
				fmt.Println("Ошибка в отправителе кошелька")
				return err
			}

			if err := s.store.Wallet().Update(ctx, toWallet); err != nil {
				fmt.Println("Ошибка в получателе кошелька")
				return err
			}

			fmt.Println("fromWallet: ", fromWallet.Balance)
			fmt.Println("toWallet: ", toWallet.Balance)

			// fmt.Println("fromWallet: ", fromWallet.Address, "переслать: ", req.Amount, "toWallet: ", toWallet.Address)
			return nil

		})

		if err != nil {
			fmt.Println("Ошибка с транзакциями")
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
