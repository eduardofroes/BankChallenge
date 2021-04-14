package transfers

import (
	"bankchallenge/commons"
	"bankchallenge/configs"
	"database/sql"
	"fmt"

	"github.com/docker/distribution/uuid"
)

// TransferRepository struct is reponsible to persist transfer data in the database.
type TransferRepository struct {
	database *sql.DB
}

var (
	transferRepository *ITransfersRepository
)

func GetTransferRepository() *ITransfersRepository {

	if transferRepository == nil {
		transferRepository = accountRepositoryBuilder()
	}

	return transferRepository
}

func accountRepositoryBuilder() *ITransfersRepository {

	transferRepository := new(TransferRepository)

	var iTransfersRepository ITransfersRepository

	iTransfersRepository = transferRepository

	return &iTransfersRepository
}

func (transferRepository *TransferRepository) openDatabase() {
	transferRepository.database = configs.GetDatabase()
}

func (transferRepository *TransferRepository) Save(transfer Transfer) {

	transferRepository.openDatabase()

	sqlStatement := "INSERT INTO transfers VALUES ($1, $2, $3, $4, $5)"

	insert, err := transferRepository.database.Prepare(sqlStatement)
	commons.CheckError(err, "Error in insert build query on transfer context.")

	result, err := insert.Exec(transfer.ID, transfer.AccountOriginId, transfer.AccountDestinationId, transfer.Amount, transfer.CreatedAt)
	commons.CheckError(err, "Error in insert transfer execution.")

	_, err = result.RowsAffected()
	commons.CheckError(err, "Error in get affected rows in transfer context.")

	defer transferRepository.database.Close()

}

func (transferRepository *TransferRepository) List() *[]Transfer {

	transferRepository.openDatabase()

	sqlStatement, err := transferRepository.database.Query("SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfers")

	commons.CheckError(err, "Error in list all tranfers query.")

	transfers := []Transfer{}

	for sqlStatement.Next() {
		var transfer Transfer

		err = sqlStatement.Scan(&transfer.ID, &transfer.AccountOriginId, &transfer.AccountDestinationId, &transfer.Amount, &transfer.CreatedAt)

		commons.CheckError(err, "Error in scan transfer data.")

		transfers = append(transfers, transfer)
	}

	defer transferRepository.database.Close()

	return &transfers
}

func (transferRepository *TransferRepository) GetTransfers(accountId uuid.UUID) *[]Transfer {

	transferRepository.openDatabase()

	sqlStatement, err := transferRepository.database.Query(fmt.Sprintf("SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfers WHERE account_origin_id = '%s' OR account_destination_id = '%s'", accountId.String(), accountId.String()))

	commons.CheckError(err, "Error in list all tranfers query.")

	transfers := []Transfer{}

	for sqlStatement.Next() {
		var transfer Transfer

		err = sqlStatement.Scan(&transfer.ID, &transfer.AccountOriginId, &transfer.AccountDestinationId, &transfer.Amount, &transfer.CreatedAt)

		commons.CheckError(err, "Error in scan transfer data.")

		transfers = append(transfers, transfer)
	}

	defer transferRepository.database.Close()

	return &transfers
}
