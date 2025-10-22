package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	controllers "github.com/shanu-shr/MagicStreamMovies/Server/MagicMovieStreamServer/Controllers"
)

func main() {
	fmt.Println("Hello world!!!")

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, magic stream movies !")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/movie", controllers.AddMovie())

	err := router.Run(":3000")
	if err != nil {
		log.Fatal("Unable to start server on port 3000")
	}
}
