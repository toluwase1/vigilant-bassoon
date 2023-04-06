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

func TestCreateUser(t *testing.T) {
	defer internal.EmptyDB()
	r := gin.Default()
	r.POST("/users/create", CreateUser)
	expectedRes := gin.H{
		"message": "user successfully created",
		"data":    gofakeit.UUID(),
	}
	testCases := []struct {
		name        string
		requestBody types.UserRequest
		assertions  func(res *httptest.ResponseRecorder, body []byte, err error)
	}{
		{
			name:        "create user successful",
			requestBody: types.UserRequest{Name: "John Doe", BVN: "1234567890"},
			assertions: func(res *httptest.ResponseRecorder, body []byte, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				var actualRes gin.H
				assert.NoError(t, json.Unmarshal(body, &actualRes))
				assert.NotEmpty(t, "user successfully created", actualRes)
			},
		},
		{
			name:        "bad request case",
			requestBody: types.UserRequest{Name: "", BVN: ""},
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

			req, err := http.NewRequest(http.MethodPost, "/users/create", bytes.NewBuffer(requestBodyBytes))
			assert.Nil(t, err)

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			tc.assertions(res, res.Body.Bytes(), err)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	defer internal.EmptyDB()
	// Initialize a test router with the GetAllUsers endpoint
	r := gin.Default()
	r.GET("/users", GetAllUsers)

	// Populate the database with some test users
	user1 := types.User{
		ID:                 "1",
		Name:               "John Doe",
		AvailableBalance:   10000,
		PendingBalance:     0,
		VerificationStatus: true,
		BVN:                "123456789",
		CreatedAt:          time.Now(),
	}
	user2 := types.User{
		ID:                 "2",
		Name:               "Jane Smith",
		AvailableBalance:   5000,
		PendingBalance:     0,
		VerificationStatus: false,
		BVN:                "987654321",
		CreatedAt:          time.Now(),
	}
	internal.SaveToDB(internal.UserTableName, user1)
	internal.SaveToDB(internal.UserTableName, user2)

	// Send a GET request to the GetAllUsers endpoint
	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code and body
	assert.Equal(t, http.StatusOK, w.Code)
	expected := gin.H{
		"message": "users successfully fetched",
		"data": []any{
			user1,
			user2,
		},
	}
	marshal, err := json.Marshal(expected)
	if err != nil {
		return
	}
	assert.Equal(t, string(marshal), w.Body.String())

}
