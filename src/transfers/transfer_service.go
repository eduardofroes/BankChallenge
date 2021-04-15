package transfers

import (
	"bankchallenge/accounts"
	"bankchallenge/commons"
	"time"

	"github.com/docker/distribution/uuid"
)

type TransferService struct {
	transferRepository *ITransfersRepository
	accountService     *accounts.AccountService
}

var (
	transferService *TransferService
)

func GetTransferService() *TransferService {

	if transferService == nil {
		transferService = transferServiceBuilder()
	}

	return transferService
}

func transferServiceBuilder() *TransferService {

	transferService = new(TransferService)

	transferService.transferRepository = GetTransferRepository()
	transferService.accountService = accounts.GetAccountService()

	return transferService
}

func (transferService *TransferService) ListAllTransfers() *[]Transfer {
	return (*transferRepository).List()
}

func (transferService *TransferService) TransferMoney(transfer Transfer) (string, bool) {

	accountOriginIdUUID, err := uuid.Parse(transfer.AccountOriginId)

	commons.CheckError(err, "Error to parse account origin id.")

	accountOrigin := transferService.accountService.GetAccount(accountOriginIdUUID)

	if accountOrigin != nil {
		if accountOrigin.Balance < transfer.Amount {
			return "FUNDS_001", false
		}
	} else {
		return "ACCOUNT_001", false
	}

	accountDestinationUUID, err := uuid.Parse(transfer.AccountDestinationId)

	commons.CheckError(err, "Error to parse account destination id.")

	accountDestination := transferService.accountService.GetAccount(accountDestinationUUID)

	if accountDestination == nil {
		return "ACCOUNT_002", false
	}

	if accountOriginIdUUID.String() != accountDestinationUUID.String() {
		accountOrigin.Balance -= transfer.Amount
		accountDestination.Balance += transfer.Amount

		transferService.accountService.UpdateAccount(*accountOrigin)
		transferService.accountService.UpdateAccount(*accountDestination)

		transfer.ID = uuid.Generate().String()
		transfer.CreatedAt = time.Now()

		(*transferRepository).Save(transfer)

		return "TRANSFER_001", true
	} else {
		return "TRANSFER_002", false
	}
}

func (transferService *TransferService) GetAccountTransfers(accountId uuid.UUID) *[]Transfer {
	return (*transferRepository).GetTransfers(accountId)
}
