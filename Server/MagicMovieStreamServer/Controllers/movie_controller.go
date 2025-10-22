package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/shanu-shr/MagicStreamMovies/Server/MagicMovieStreamServer/Database"
	models "github.com/shanu-shr/MagicStreamMovies/Server/MagicMovieStreamServer/Models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movie")

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.JSON(200, gin.H{"message": "list of movies"})

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies []models.Movie

		cursor, err := movieCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		}

		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
		}

		c.JSON(200, movies)
	}
}
