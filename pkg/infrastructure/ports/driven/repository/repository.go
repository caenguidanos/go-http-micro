package repository

import (
	"time"

	"microservice/pkg/domain"
	"microservice/pkg/infrastructure/adapters"
)

type Repository struct {
	db adapters.DatabaseAdapter
}

func NewRepository(db adapters.DatabaseAdapter) *Repository {
	return &Repository{db: db}
}

func (r Repository) InsertOne(user string, email string) (string, *domain.HttpException) {
	doc := domain.Stuff{
		User:      user,
		Email:     email,
		CreatedAt: uint32(time.Now().UnixNano()),
		UpdatedAt: uint32(time.Now().UnixNano()),
	}

	return r.db.InsertOne(doc)
}
