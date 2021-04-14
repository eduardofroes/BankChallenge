package login

import (
	"bankchallenge/accounts"
	"bankchallenge/commons"
)

// LoginService struct is responsable to apply and check the user credentials.
type LoginService struct {
	accountService    *accounts.AccountService
	accountRepository *accounts.IAccountRepository
}

var (
	loginService *LoginService
)

func GetLoginService() *LoginService {

	if loginService == nil {
		loginService = loginServiceBuilder()
	}

	return loginService
}

func loginServiceBuilder() *LoginService {

	loginService = new(LoginService)

	loginService.accountService = accounts.GetAccountService()

	return loginService
}

func (loginService *LoginService) Login(login Login) *string {

	account := loginService.accountService.GetCheckByCredentials(login.CPF, login.Secret)

	if account != nil {

		token, err := commons.GenerateToken(login.CPF, login.Secret)

		commons.CheckError(err, "Error in generates the login token.")

		return token

	}

	return nil
}
