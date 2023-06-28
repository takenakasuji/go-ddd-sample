package service

import (
	"github.com/takenakasuji/go-ddd-sample/internal/repository"
)

type ChefService struct {
	ChefRepository repository.ChefRepository
}

func NewChefService(repository repository.ChefRepository) *ChefService {
	return &ChefService{
		ChefRepository: repository,
	}
}

func (cs *ChefService) Exist(name string) bool {
	c := cs.ChefRepository.ByName(name)
	if c.Name == "" {
		return false
	}
	return true
}
