package types

type UserRequest struct {
	Name string `json:"name"`
	BVN  string `json:"bvn"`
}

type TransactionRequest struct {
	FromId string `json:"from_id"`
	ToId   string `json:"to_id"`
	Amount int64  `json:"amount"`
}
type TopUpRequest struct {
	UserId string `json:"user_id"`
	Amount int64  `json:"amount"`
}
