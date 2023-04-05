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

// i get this error, please fix it: Error:      	Expected nil, but got: &json.SyntaxError{msg:"invalid character 'p' after top-level value", Offset:5}
var router = gin.Default()

type Data struct {
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
	Errors    string      `json:"errors,omitempty"`
	Status    string      `json:"status,omitempty"`
}

//func TestCreateUser(t *testing.T) {
//	// create a test gin context
//	//r := gin.Default()
//	w := httptest.NewRecorder()
//	c, _ := gin.CreateTestContext(w)
//
//	// create a user request body
//	userReq := UserRequest{
//		Name: "John Doe",
//		BVN:  "1234567890",
//	}
//
//	// bind the user request body to the gin context
//	c.Request, _ = http.NewRequest("POST", "/users/create", nil)
//	c.Request.Header.Set("Content-Type", "application/json")
//	reqBody, _ := json.Marshal(userReq)
//	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
//
//	// call the CreateUser function
//	CreateUser(c)
//
//	// check the response
//	assert.Equal(t, http.StatusOK, w.Code)
//	var res map[string]interface{}
//	json.Unmarshal(w.Body.Bytes(), &res)
//	assert.Equal(t, "user successfully created", res["message"])
//	assert.NotNil(t, res["data"])
//
//}

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
