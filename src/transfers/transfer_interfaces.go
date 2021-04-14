package transfers

import (
	"github.com/docker/distribution/uuid"
)

// IAccountRepository is an interface used to account repository implementation.
type ITransfersRepository interface {
	Save(transfer Transfer)
	List() *[]Transfer
	GetTransfers(accountId uuid.UUID) *[]Transfer
}
