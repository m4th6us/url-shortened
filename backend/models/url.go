package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID 			primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Original 	string `bson:"original" json:"original"`
	ShortCode 	string `bson:"short_code" json:"short_code"`
}