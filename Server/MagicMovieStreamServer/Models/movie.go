package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Movie struct {
	ID          bson.ObjectID
	ImdbID      string
	Title       string
	PosterPath  string
	YouTubeId   string
	Genre       []Genre
	AdminReview string
	Ranking     []Ranking
}
