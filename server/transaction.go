package server

import (
	"github.com/gin-gonic/gin"
	"lemonadee/internal"
	"lemonadee/render"
	"lemonadee/types"
	"net/http"
)

func CreateTransaction(c *gin.Context) {
	request := types.TransactionRequest{}
	if err := c.BindJSON(request); err != nil {
		render.BadRequest(c)
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
	id, err := internal.SaveToDB(internal.UserTableName, transaction)
	if err != nil {
		return "", internal.CoverError("server/createTransaction", err)
	}
	internal.PushToQueue(internal.UserQueue, transaction)
	return id, nil
}
