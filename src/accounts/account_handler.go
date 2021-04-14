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

	if commons.ValidateToken(r.Header.Get("authorization")) {
		GetAccountService()

		accounts := accountService.ListAllAccounts()

		if accounts == nil {
			commons.HandleNotFound(w, "No accounts was registred.")
			return
		}

		commons.WriteJSON(w, accounts, 200)
	} else {
		commons.HandleUnauthorized(w, "Resource unauthorized.")
	}
}

// GetAccountBalanceHandler get the balance from an account.
func GetAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	defer commons.HandleError(w)

	if commons.ValidateToken(r.Header.Get("authorization")) {
		GetAccountService()

		vars := mux.Vars(r)
		accountID := vars["id"]

		accountUUID, err := uuid.Parse(accountID)

		if err != nil {
			commons.HandleBadRequest(w, "Error to parse account id.")
			return
		}

		account := accountService.GetAccount(accountUUID)

		if account == nil {
			commons.HandleNotFound(w, fmt.Sprintf("Account id %v not found", accountID))
			return
		}

		body := map[string]float64{
			"balance": account.Balance,
		}

		commons.WriteJSON(w, body, 200)
	} else {
		commons.HandleUnauthorized(w, "Resource unauthorized.")
	}

}

// CreateAccountHandler is responsible for create user account.
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	defer commons.HandleError(w)

	GetAccountService()

	var account Account

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		commons.HandleBadRequest(w, "Could not convert the account body.")
		return
	}

	secret := accountService.CreateAccount(&account)

	body := map[string]string{
		"secret": secret,
	}

	commons.WriteJSON(w, body, 201)
}
