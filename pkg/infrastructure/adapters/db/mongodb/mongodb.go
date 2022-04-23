package mongodb

import (
	"context"
	"net/http"
	"time"

	"ecommerce-template/apps/microservice/auth/credential/pkg/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbAdapter struct {
	conn *mongo.Client
}

func NewAdapter(conn *mongo.Client) *mongodbAdapter {
	return &mongodbAdapter{conn: conn}
}

func (m mongodbAdapter) InsertOne(doc domain.Stuff) (string, *domain.HttpException) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var fallback string

	rs, e := m.conn.Database("service").Collection("stuff").InsertOne(ctx, doc)
	if e != nil {
		return fallback, &domain.HttpException{
			Status:      http.StatusConflict,
			Code:        "#",
			Description: http.StatusText(http.StatusConflict),
		}
	}

	return rs.InsertedID.(primitive.ObjectID).Hex(), nil
}
