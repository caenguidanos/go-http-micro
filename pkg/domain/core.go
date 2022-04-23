package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stuff struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User      string             `json:"user" bson:"user"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt uint32             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt uint32             `json:"updated_at" bson:"updated_at,omitempty"`
}

type HttpException struct {
	Status      int
	Code        string
	Description string
}
