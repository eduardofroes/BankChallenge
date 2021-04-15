package accounts

import (
	"bankchallenge/commons"
	"bankchallenge/configs"
	"database/sql"
	"fmt"

	"github.com/docker/distribution/uuid"
)

// AccountRepository struct is reponsible to persist account data in the database.
type AccountRepository struct {
	database *sql.DB
}

var (
	accountRepository *IAccountRepository
)

func GetAccountRepository() *IAccountRepository {

	if accountRepository == nil {
		accountRepository = accountRepositoryBuilder()
	}

	return accountRepository
}

func accountRepositoryBuilder() *IAccountRepository {

	accountRepository := new(AccountRepository)

	var iAccountRepository IAccountRepository

	iAccountRepository = accountRepository

	return &iAccountRepository
}

func (accountRepository *AccountRepository) openDatabase() {
	accountRepository.database = configs.GetDatabase()
}

func (accountRepository *AccountRepository) List() *[]Account {

	accountRepository.openDatabase()

	sqlStatement, err := accountRepository.database.Query("SELECT id, name, cpf, secret, balance, created_at FROM accounts")

	commons.CheckError(err, "Error in list all accounts query.")

	accounts := []Account{}

	for sqlStatement.Next() {
		var account Account

		err = sqlStatement.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)

		commons.CheckError(err, "Error in scan account data.")

		accounts = append(accounts, account)
	}

	defer accountRepository.database.Close()

	return &accounts
}

func (accountRepository *AccountRepository) Get(id uuid.UUID) *Account {

	accountRepository.openDatabase()

	var account Account

	err := accountRepository.database.QueryRow(fmt.Sprintf("SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE id = '%s'", id.String())).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)

	defer accountRepository.database.Close()

	if err != nil {
		return nil
	} else {
		return &account
	}
}

func (accountRepository *AccountRepository) GetCheckByCredentials(CPF string, secret string) *Account {

	accountRepository.openDatabase()

	var account Account

	err := accountRepository.database.QueryRow(fmt.Sprintf("SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE cpf ilike '%s' and secret ilike '%s'", CPF, secret)).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)

	defer accountRepository.database.Close()

	if err != nil {
		return nil
	} else {
		return &account
	}
}

func (accountRepository *AccountRepository) Save(account Account) {

	accountRepository.openDatabase()

	sqlStatement := "INSERT INTO accounts VALUES ($1, $2, $3, $4, $5, $6)"

	insert, err := accountRepository.database.Prepare(sqlStatement)
	commons.CheckError(err, "Error in insert build query on account context.")

	result, err := insert.Exec(account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)
	commons.CheckError(err, "Error in insert account execution.")

	_, err = result.RowsAffected()
	commons.CheckError(err, "Error in get affected rows in account context.")

	defer accountRepository.database.Close()

}

func (accountRepository *AccountRepository) Update(account Account) {

	accountRepository.openDatabase()

	sqlStatement := "UPDATE accounts SET name=$1, cpf=$2, secret=$3, balance=$4 WHERE id = $5"

	update, err := accountRepository.database.Prepare(sqlStatement)
	commons.CheckError(err, "Error in update build query on account context.")

	result, err := update.Exec(account.Name, account.CPF, account.Secret, account.Balance, account.ID)
	commons.CheckError(err, "Error in update account execution.")

	_, err = result.RowsAffected()
	commons.CheckError(err, "Error in get affected rows in account context.")

	defer accountRepository.database.Close()

}
