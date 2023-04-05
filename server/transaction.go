package server

import (
	"github.com/gin-gonic/gin"
	"lemonadee/internal"
	"lemonadee/render"
	"lemonadee/types"
	"log"
	"net/http"
	"time"
)

func CreateTransaction(c *gin.Context) {
	request := types.TransactionRequest{}
	if err := c.BindJSON(&request); err != nil {
		render.BadRequest(c)
		return
	}
	_, err := internal.GetByID(internal.UserTableName, request.ToId)
	if err != nil {
		render.Error(c, err)
		return
	}

	id, err := createTransaction(request)
	if err != nil {
		render.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "transaction successfully created",
		"data":    id,
	})
}

func createTransaction(request types.TransactionRequest) (string, *internal.Error) {
	transaction := types.NewTransaction(request.FromId, request.ToId, request.Amount)
	id, err := internal.SaveToDB(internal.TransactionTableName, transaction)
	if err != nil {
		return "", internal.CoverError("server/createTransaction", err)
	}
	data, err := internal.GetByID(internal.UserTableName, transaction.FromId)
	if err != nil {
		return "", internal.CoverError("server/createTransaction", err)
	}
	user := types.UserFromDB(data)
	if user.Balance < request.Amount {
		return "", internal.NewError("insufficient balance", http.StatusBadRequest)
	}
	if user.VerificationStatus {
		internal.PushToQueue(internal.TransactionQueue, transaction)
		transaction.Status = "pushed-to-queue"
		transaction.UpdatedAt = time.Now()
		err := internal.UpdateDB(internal.TransactionTableName, transaction)
		if err != nil {
			return "", internal.CoverError("server/createTransaction", err)
		}
	} else {
		internal.PushToQueue(internal.UserQueue, user)
	}
	return id, nil
}

func processTransaction(transaction types.Transactions) *internal.Error {
	err := internal.UpdateDB(internal.TransactionTableName, transaction)
	if err != nil {
		return internal.CoverError("server/createTransaction", err)
	}
	fromUser, err := internal.GetByID(internal.UserTableName, transaction.FromId)
	if err != nil {
		return internal.CoverError("server/createTransaction", err)
	}
	toUser, err := internal.GetByID(internal.UserTableName, transaction.ToId)
	if err != nil {
		return internal.CoverError("server/createTransaction", err)
	}
	fromUsr := types.UserFromDB(fromUser)
	toUsr := types.UserFromDB(toUser)
	f := func() []any {
		transaction.Status = "processed"
		transaction.UpdatedAt = time.Now()
		fromUsr.Balance -= transaction.Amount
		toUsr.Balance += transaction.Amount
		return []any{fromUsr, toUsr}
	}
	return internal.UpdateDbTx(f, internal.UserTableName)
}

func transactionConsumer() {
	data := internal.GetAllFromQueue(internal.TransactionQueue)
	for _, v := range data {
		if err := processTransaction(types.TransactionFromDB(v)); err != nil {
			log.Println("error while processing transaction: ", err)
			internal.PushToQueue(internal.TransactionQueue, v)
		}
	}
}
