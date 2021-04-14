package accounts

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/docker/distribution/uuid"
)

// AccountService struct is responsable to apply business rules related to user account.
type AccountService struct {
	accountRepository *IAccountRepository
}

var (
	accountService *AccountService
)

func GetAccountService() *AccountService {

	if accountService == nil {
		accountService = accountServiceBuilder()
	}

	return accountService
}

func accountServiceBuilder() *AccountService {

	accountService = new(AccountService)

	accountService.accountRepository = GetAccountRepository()

	return accountService
}

func (accountService *AccountService) ListAllAccounts() *[]Account {
	return (*accountRepository).List()
}

func (accountService *AccountService) GetAccount(id uuid.UUID) *Account {
	return (*accountRepository).Get(id)
}

func (accountService *AccountService) GetCheckByCredentials(CPF string, secret string) *Account {
	return (*accountRepository).GetCheckByCredentials(CPF, secret)
}

func (accountService *AccountService) UpdateAccount(account Account) {
	(*accountRepository).Update(account)
}

func (accountService *AccountService) CreateAccount(account *Account) string {

	account.ID = uuid.Generate().String()

	if account.Balance < 0 {
		account.Balance = 0
	}

	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%s.%s", account.Name, account.CPF)))
	secret := hex.EncodeToString(hash.Sum(nil))

	account.Secret = secret
	account.CreatedAt = time.Now()

	(*accountRepository).Save(*account)

	return secret
}
