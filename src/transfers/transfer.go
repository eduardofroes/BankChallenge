package transfers

import (
	"time"
)

// Transfer data struct that represents the user transfer.
type Transfer struct {
	ID                   string    `json:"id"`
	AccountOriginId      string    `json:"account_origin_id"`
	AccountDestinationId string    `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}
