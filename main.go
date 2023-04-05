package main

import (
	"github.com/gin-gonic/gin"
	"lemonadee/server"
)

/*
Create an endpoint that creates users with IDs as counts (i.e  user1 :1  user2 :2 user3 :3).
(ID, Name, balance and Verification_Status)

Users should go through verification & should be put in a verification queue
Verification should be processed periodically by X amount of workers (goroutines) that are spun up in X amount of time e.g every 30s
Create an endpoint that creates transactions taking in userID & amount to send
The transactions should be put in a queue that will be processed
Transaction should be processed periodically by X amount of workers (goroutines) that are spun up in X amount of time e.g every 30s
Transactions should only be processed for verified users - If user is unverified, user should be pushed to the verification queue and verified
Create an endpoint that returns all users and their balances and verification status
BONUS: Figure out a way that user will not be verified even if pushed to the verification queue.

*/

func main() {
	r := gin.Default()
	r.POST("/users/create", server.CreateUser)
	r.POST("/transactions/create", server.CreateTransaction)
	r.GET("/users", server.GetAllUsers)
	go server.Consume()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
