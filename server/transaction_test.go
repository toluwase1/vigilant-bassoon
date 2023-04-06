package server

import (
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"lemonadee/internal"
	"lemonadee/types"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateTransactions(t *testing.T) {
	defer internal.EmptyDB()
	r := gin.Default()
	r.POST("/transactions/create", CreateUser)
	expectedRes := gin.H{
		"message": "user successfully created",
		"data":    gofakeit.UUID(),
	}
	testCases := []struct {
		name        string
		requestBody types.TransactionRequest
		assertions  func(res *httptest.ResponseRecorder, body []byte, err error)
	}{
		{
			name:        "bad request case",
			requestBody: types.TransactionRequest{FromId: "John Doe", ToId: "1234567890", Amount: 1000},
			assertions: func(res *httptest.ResponseRecorder, body []byte, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusBadRequest, res.Code)
				var actualRes gin.H
				assert.NoError(t, json.Unmarshal(body, &actualRes))
				assert.NotEmpty(t, expectedRes, actualRes)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, err := json.Marshal(tc.requestBody)
			assert.Nil(t, err)

			req, err := http.NewRequest(http.MethodPost, "/transactions/create", bytes.NewBuffer(requestBodyBytes))
			assert.Nil(t, err)

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			tc.assertions(res, res.Body.Bytes(), err)
		})
	}
}

func loadDB() {
	user1 := types.User{ID: "1", Name: "John Doe", AvailableBalance: 10000, VerificationStatus: true, BVN: "123456789", CreatedAt: time.Now()}
	user2 := types.User{ID: "2", Name: "Jane Smith", AvailableBalance: 5000, VerificationStatus: false, BVN: "987654321", CreatedAt: time.Now()}
	transaction1 := types.Transactions{ID: "1", FromId: user1.ID, ToId: user2.ID, Amount: 500, Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	transaction2 := types.Transactions{ID: "2", FromId: user2.ID, ToId: user1.ID, Amount: 200, Status: "completed", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	internal.SaveToDB(internal.UserTableName, user1)
	internal.SaveToDB(internal.UserTableName, user2)
	internal.SaveToDB(internal.TransactionTableName, transaction1)
	internal.SaveToDB(internal.TransactionTableName, transaction2)
}
