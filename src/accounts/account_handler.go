package accounts

import (
	"bankchallenge/commons"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/distribution/uuid"
	"github.com/gorilla/mux"
)

// ListAccountsHandler list all accounts registred
func ListAccountsHandler(w http.ResponseWriter, r *http.Request) {

	defer commons.HandleError(w)

	GetAccountService()

	accounts := accountService.ListAllAccounts()

	if accounts == nil {
		commons.HandleNotFound(w, "No accounts was registred.")
		return
	}

	commons.WriteJSON(w, accounts, 200)
}

// GetAccountBalanceHandler get the balance from an account.
func GetAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	defer commons.HandleError(w)

	GetAccountService()

	vars := mux.Vars(r)
	accountID := vars["id"]

	accountUUID, _ := uuid.Parse(accountID)

	account := accountService.GetAccount(accountUUID)

	if account == nil {
		commons.HandleNotFound(w, fmt.Sprintf("Account id %v not found", accountID))
		return
	}

	body := map[string]float64{
		"balance": account.Balance,
	}

	commons.WriteJSON(w, body, 200)
}

// CreateAccountHandler is responsible for create user account.
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	defer commons.HandleError(w)

	GetAccountService()

	var account Account

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		panic("Could not convert the account body.")
	}

	secret := accountService.CreateAccount(account)

	body := map[string]string{
		"secret": secret,
	}

	commons.WriteJSON(w, body, 201)
}
