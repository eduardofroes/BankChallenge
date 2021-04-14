package login

import (
	"bankchallenge/commons"
	"encoding/json"
	"net/http"
)

// LoginHandler checks and generates the login credentials.
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	defer commons.HandleError(w)

	GetLoginService()

	var login Login

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		commons.HandleBadRequest(w, "Could not convert the login body.")
		return
	}

	token := loginService.Login(login)

	body := map[string]string{
		"token": *token,
	}

	commons.WriteJSON(w, body, 200)
}
