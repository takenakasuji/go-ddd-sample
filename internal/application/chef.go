package application

import (
	"fmt"
	"strconv"

	"github.com/takenakasuji/go-ddd-sample/internal/domain/service"
	"github.com/takenakasuji/go-ddd-sample/internal/dto"
	"github.com/takenakasuji/go-ddd-sample/internal/repository"
)

type chefApplicationService struct {
	chefRepository repository.ChefRepository
}

type ChefApplicationService interface {
	List() *dto.Chefs
	ChefProfile(sid string) *dto.Chef
	CreateChef(name string) error
}

func NewChefApplicationService(repository repository.ChefRepository) ChefApplicationService {
	return &chefApplicationService{
		chefRepository: repository,
	}
}

func (app *chefApplicationService) List() *dto.Chefs {
	return app.chefRepository.List()
}

func (app *chefApplicationService) ChefProfile(sid string) *dto.Chef {
	id, _ := strconv.Atoi(sid)
	return app.chefRepository.ByID(id)
}

func (app *chefApplicationService) CreateChef(name string) error {
	s := service.NewChefService(app.chefRepository)
	// シェフが存在するかはエンティティの立ちふるまいではないためドメインサービスとして処理する
	if s.Exist(name) {
		return fmt.Errorf("the name is existed")
	}
	return app.chefRepository.Create(name)
}
