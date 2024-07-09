package main

import (
	// "fmt"
	"pdfGen/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/createMemarts", controllers.CreateMemarts)

	err := r.Run()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
