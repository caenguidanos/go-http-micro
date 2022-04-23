package main

import (
	"context"
	"os"

	"microservice/pkg/application/core/service"
	"microservice/pkg/infrastructure/adapters/db/mongodb"
	"microservice/pkg/infrastructure/ports/driven/repository"
	"microservice/pkg/infrastructure/ports/driver/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := mongodb.Connect(os.Getenv("MONGODB_URI"))
	defer conn.Disconnect(context.Background())

	controller := handler.New(
		service.NewService(
			repository.NewRepository(
				mongodb.NewAdapter(conn),
			),
		),
	)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.Default())

	group := r.Group("/v1")
	{
		group.POST("/stuff", controller.CreateStuff)
	}

	r.Run()
}
