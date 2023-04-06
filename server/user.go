package server

import (
	"github.com/gin-gonic/gin"
	"lemonadee/internal"
	"lemonadee/render"
	"lemonadee/types"
	"log"
	"net/http"
	"strings"
)

func CreateUser(c *gin.Context) {
	userR := types.UserRequest{}
	if err := c.BindJSON(&userR); err != nil {
		log.Println("error", err)
		render.BadRequest(c)
		return
	}
	_, err := validateRequest(userR.Name)
	if err != nil {
		render.Error(c, err)
		return
	}

	userID, err := createUser(userR)
	if err != nil {
		render.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully created",
		"data":    userID,
	})

}

func GetAllUsers(c *gin.Context) {
	users, err := internal.GetAllFromDB(internal.UserTableName)
	if err != nil {
		render.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "users successfully fetched",
		"data":    users,
	})

}

func createUser(request types.UserRequest) (string, *internal.Error) {
	usr := types.NewUser(request.Name, request.BVN)
	userID, err := internal.SaveToDB(internal.UserTableName, usr)
	if err != nil {
		return "", internal.CoverError("server/createUser", err)
	}
	internal.PushToQueue(internal.UserQueue, usr)
	return userID, nil
}

func verifyUser(user types.User) *internal.Error {
	if user.BVN != "" {
		user.VerificationStatus = true
		err := internal.UpdateDB(internal.UserTableName, user)
		if err != nil {
			return internal.CoverError("server/verifyUser", err)
		}
	}
	return nil
}

func userConsumer() {
	data := internal.GetAllFromQueue(internal.UserQueue)
	for _, v := range data {
		if err := verifyUser(types.UserFromDB(v)); err != nil {
			log.Println("error while processing transaction: ", err)
			internal.PushToQueue(internal.UserQueue, v)
		}
	}
}

func validateRequest(name string) (string, *internal.Error) {
	err := &internal.Error{
		Message:    "name field cannot be blank",
		StatusCode: http.StatusBadRequest,
	}
	if strings.TrimSpace(name) != "" {
		return name, nil
	}
	return "", err
}
