package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world!!")

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, magic stream movies !")
	})

	err := router.Run(":3000")
	if err != nil {
		log.Fatal("Unable to start server on port 3000")
	}

}
