package accounts

import (
	"github.com/docker/distribution/uuid"
)

// IAccountRepository is an interface used to account repository implementation.
type IAccountRepository interface {
	Save(account Account)
	List() *[]Account
	Get(id uuid.UUID) *Account
	Update(account Account)
}
