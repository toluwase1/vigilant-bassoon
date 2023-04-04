package types

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Balance            int64     `json:"balance"`
	VerificationStatus bool      `json:"verification_status"`
	BVN                string    `json:"bvn"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewUser(name, bvn string) *User {
	return &User{
		ID:                 uuid.New().String(),
		Name:               name,
		Balance:            0,
		VerificationStatus: false,
		BVN:                bvn,
		CreatedAt:          time.Now(),
	}
}

type Transactions struct {
	ID        string    `json:"id"`
	FromId    string    `json:"from_id"`
	ToId      string    `json:"to_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTransaction(fromId, ToId string, Amount int64) *Transactions {
	return &Transactions{
		ID:        uuid.New().String(),
		FromId:    fromId,
		ToId:      ToId,
		Amount:    Amount,
		CreatedAt: time.Now(),
	}
}
