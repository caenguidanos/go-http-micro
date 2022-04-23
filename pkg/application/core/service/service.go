package service

import (
	"microservice/pkg/domain"
	"microservice/pkg/infrastructure/ports"
)

type service struct {
	repo ports.RepositoryPort
}

func NewService(repo ports.RepositoryPort) *service {
	return &service{repo: repo}
}

func (s service) CreateStuff() (string, *domain.HttpException) {
	return s.repo.InsertOne("caenguidanos", "1234")
}
