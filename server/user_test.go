package server

import (
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"lemonadee/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Data struct {
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
	Errors    string      `json:"errors,omitempty"`
	Status    string      `json:"status,omitempty"`
}

func TestCreateUser(t *testing.T) {
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

func TopupTest(t *testing.T) {
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
