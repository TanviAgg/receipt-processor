package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"receipt-processor/internal/handler"
)

func main() {
	router := gin.Default()

	router.GET("/receipts/:id/points", handler.GetPoints)
	router.POST("/receipts/process", handler.ProcessReceipts)

	log.Println("Listening on localhost:8080")
	router.Run()
}
