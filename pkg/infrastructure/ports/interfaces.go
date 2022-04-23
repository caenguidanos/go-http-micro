package ports

import (
	"microservice/pkg/domain"

	"github.com/gin-gonic/gin"
)

type ServicePort interface {
	CreateStuff() (string, *domain.HttpException)
}

type RepositoryPort interface {
	InsertOne(user string, email string) (string, *domain.HttpException)
}

type HttpHandlerPort interface {
	CreateStuff(c *gin.Context)
}
