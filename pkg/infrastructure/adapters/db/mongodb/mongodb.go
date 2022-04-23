package mongodb

import (
	"context"
	"log"
	"net/http"
	"time"

	"microservice/pkg/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) *mongo.Client {
	c, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)),
	)

	if err != nil {
		log.Fatal(err)
	}

	return c
}

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
