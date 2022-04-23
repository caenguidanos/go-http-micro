package adapters

import "microservice/pkg/domain"

type DatabaseAdapter interface {
	InsertOne(doc domain.Stuff) (string, *domain.HttpException)
}
