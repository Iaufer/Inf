package apiserver

import (
	"Inf/internal/app/apiserver/store"
	"Inf/internal/app/model"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	s.router.HandleFunc("/api/transactions", s.handleGetTransactions()).Methods("GET")
	s.router.Handle("/api/wallet/{address}/balance", s.handleGetBalance()).Methods("GET")

	s.router.HandleFunc("/api/wallet", s.handleCreateWallet()).Methods("POST")
}

func (s *server) handleCreateWallet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes := make([]byte, 32)

		if _, err := rand.Read(bytes); err != nil {
			http.Error(w, fmt.Sprintf("failed to generate random bytes: %v", err), http.StatusInternalServerError)
			return
		}

		addr := hex.EncodeToString(bytes)

		wallet := &model.Wallet{
			Address: addr,
			Balance: 0,
		}

		ctx := r.Context()

		if err := s.store.Wallet().Create(ctx, wallet); err != nil {
			http.Error(w, "Failed to create wallet", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(wallet); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handleGetBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		ctx := r.Context()

		wallet, err := s.store.Wallet().FindByAddress(ctx, address[1:len(address)-1])

		if err != nil {
			http.Error(w, "Wallet not found", http.StatusNotFound)
			return
		}

		response := model.Wallet{
			ID:      wallet.ID,
			Address: wallet.Address,
			Balance: wallet.Balance,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handleGetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validate := validator.New()

		countStr := r.URL.Query().Get("count")

		if countStr == "" {
			countStr = "10"
		}

		count, err := strconv.Atoi(countStr)

		if err != nil {
			http.Error(w, "Invalid count parameter", http.StatusBadRequest)
			return
		}

		req := struct {
			Count int `validate:"gte=1,lte=100"`
		}{Count: count}

		if err := validate.Struct(req); err != nil {
			http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
			return
		}

		var trs []*model.Transaction

		if err := s.store.Transaction().GetTr(req.Count, &trs); err != nil {
			http.Error(w, "Failed to fatch tr", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(trs); err != nil {
			http.Error(w, "Failed to encode transactions", http.StatusInternalServerError)
			return
		}

	}
}

func (s *server) handleSend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			From   string  `json:"from" validate:"required"`
			To     string  `json:"to" validate:"required"`
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

		if req.From == req.To {
			http.Error(w, "sent to myself", http.StatusBadRequest)
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

			tr := &model.Transaction{
				From:   req.From,
				To:     req.To,
				Amount: float64(req.Amount),
			}

			if err := s.store.Transaction().Create(ctx, tr); err != nil {
				fmt.Println("Error creating transaction records")
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
