package routes

import (
	"bankchallenge/accounts"
	"bankchallenge/login"
	"bankchallenge/transfers"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []route

var routes = Routes{
	route{
		"List all accounts registred.",
		"GET",
		"/accounts",
		accounts.ListAccountsHandler,
	},
	route{
		"Get the account balance according to id.",
		"GET",
		"/accounts/{id}/balance",
		accounts.GetAccountBalanceHandler,
	},
	route{
		"Create a new account.",
		"POST",
		"/accounts",
		accounts.CreateAccountHandler,
	},
	route{
		"Login",
		"POST",
		"/login",
		login.LoginHandler,
	},
	route{
		"Tranfer money to another account.",
		"POST",
		"/transfers",
		transfers.TransferHandler,
	},
	route{
		"List tranfers of an account.",
		"GET",
		"/transfers",
		transfers.ListTransfersHandler,
	},
}
