package transfers

import (
	"bankchallenge/commons"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/distribution/uuid"
)

// ListTransfersHandler list all transfers from one specific account
func ListTransfersHandler(w http.ResponseWriter, r *http.Request) {

	defer commons.HandleError(w)

	if commons.ValidateToken(r.Header.Get("authorization")) {
		GetTransferService()

		accountID := commons.ExtractMetadata(r.Header.Get("authorization"))

		accountUUID, err := uuid.Parse(*accountID)

		if err != nil {
			commons.HandleBadRequest(w, "Error to parse account id.")
			return
		}

		transfers := transferService.GetAccountTransfers(accountUUID)

		if transfers == nil {
			commons.HandleNotFound(w, "No transfers was registred.")
			return
		}

		commons.WriteJSON(w, transfers, 200)
	} else {
		commons.HandleUnauthorized(w, "Resource unauthorized.")
	}
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {

	defer commons.HandleError(w)

	if commons.ValidateToken(r.Header.Get("authorization")) {
		GetTransferService()

		accountID := commons.ExtractMetadata(r.Header.Get("authorization"))

		var transfer Transfer

		err := json.NewDecoder(r.Body).Decode(&transfer)

		if err != nil {
			commons.HandleBadRequest(w, "Could not convert the transfer body.")
			return
		}

		transfer.AccountOriginId = *accountID

		code, transferred := transferService.TransferMoney(transfer)

		if !transferred {

			body := map[string]string{
				"message": fmt.Sprintf("Transferred to account id: %s.", transfer.AccountDestinationId),
			}

			commons.WriteJSON(w, body, 200)

		} else {
			commons.TreatCode(w, code)
		}

	} else {
		commons.HandleUnauthorized(w, "Resource unauthorized.")
	}
}
