package uuid

import "github.com/google/uuid"

type Service interface {
	Generate() uuid.UUID
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s service) Generate() uuid.UUID {
	return uuid.New()
}
