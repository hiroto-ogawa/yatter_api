package object

import (
	"time"
)

type Status struct {
	// The internal ID of the status
	ID int64 `json:"id,omitempty"`

	// The account_id of the status
	AccountID int64 `json:"account_id,omitempty"`

	// The username of the account
	Content string `json:"status,omitempty" db:"content"`

	// The time the status was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
