package types

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	AvailableBalance   int64     `json:"available_balance"`
	PendingBalance     int64     `json:"pending_balance"`
	VerificationStatus bool      `json:"verification_status"`
	BVN                string    `json:"bvn"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewUser(name, bvn string) User {
	return User{
		ID:                 uuid.New().String(),
		Name:               name,
		AvailableBalance:   1000, //default balance
		PendingBalance:     0,
		VerificationStatus: false,
		BVN:                bvn,
		CreatedAt:          time.Now(),
	}
}
func UserFromDB(data any) User {
	user := data.(User)
	return user
}

type Transactions struct {
	ID        string    `json:"id"`
	FromId    string    `json:"from_id"`
	ToId      string    `json:"to_id"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func TransactionFromDB(data any) Transactions {
	transact := data.(Transactions)
	return transact
}

func NewTransaction(fromId, ToId string, Amount int64) Transactions {
	return Transactions{
		ID:        uuid.New().String(),
		FromId:    fromId,
		ToId:      ToId,
		Amount:    Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
