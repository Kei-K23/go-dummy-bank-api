package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/services"
)


func CreateAccountHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
		account := struct {
			UserId string `json:"user_id"`
		} {	}

		err := json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resAcc, err := services.CreateAccount(db , account.UserId)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resAcc)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
     }
}

func UpdateAccountHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        id := r.PathValue("id")

        if id == "" {
            newErr := errors.New("id is required")
            api.ErrBadRequest(w, newErr)
			return
        }

        acc := api.Account{}

		err := json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resAcc, err := services.UpdateAccount(db , acc.Balance, id)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resAcc)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
     }
}

func DepositAccountHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        id := r.PathValue("id")

        if id == "" {
            newErr := errors.New("id is required")
            api.ErrBadRequest(w, newErr)
			return
        }

        acc := api.Account{}

		err := json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resAcc, err := services.DepositAccount(db , acc.Balance, id)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resAcc)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
     }
}

func WithdrawAccountHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        id := r.PathValue("id")

        if id == "" {
            newErr := errors.New("id is required")
            api.ErrBadRequest(w, newErr)
			return
        }

        acc := api.Account{}

		err := json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resAcc, err := services.WithdrawAccount(db , acc.Balance, id)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resAcc)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
     }
}

func TransferBalanceHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        fromId := r.PathValue("fromId")
        toId := r.PathValue("toId")

        if fromId == "" {
            newErr := errors.New("fromId is required")
            api.ErrBadRequest(w, newErr)
			return
        }
        
		if toId == "" {
            newErr := errors.New("toId is required")
            api.ErrBadRequest(w, newErr)
			return
        }

        acc := api.Account{}

		err := json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resAcc, err := services.TransferBalance(db , acc.Balance, fromId, toId)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resAcc)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
     }
}

func DeleteAccountHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        id := r.PathValue("id")

        if id == "" {
            newErr := errors.New("id is required")
            api.ErrBadRequest(w, newErr)
			return
        }

		resUser, err := services.DeleteAccount(db, id)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resUser)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
     }
}

func GetAccountHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")

        authHeader := strings.Split(r.Header.Get("Authorization"), " ")[1] 
        
        if id == "" { 
            newErr := errors.New("id parameter is required")
            api.ErrBadRequest(w, newErr)
            return
        }

        acc := api.ResForLogin{
            Id:          id,
            AccessToken: authHeader,
        }
    
        resAcc, err := services.GetAccount(db, &acc)

        if err != nil {
            api.ErrInternalServer(w, err)
            return
        }

        resJson, err := json.Marshal(resAcc)

        if err != nil {
            api.ErrInternalServer(w, err)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        _, err = w.Write(resJson)
        if err != nil { 
            api.ErrInternalServer(w, err)
            return
        }
    }
}
