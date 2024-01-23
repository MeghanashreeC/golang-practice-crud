package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Netflix - it is a collection where i can watch series and movies and I can mark them watched or unwatched
type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
