package server

import (
	"log"
	"time"
)

func Consume() {

	func() {
		for {
			log.Println("cronjob running")
			go userConsumer()
			go transactionConsumer()
			time.Sleep(30 * time.Second)
		}
	}()
	select {}
}
