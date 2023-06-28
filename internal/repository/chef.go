package repository

import (
	"gorm.io/gorm"

	"github.com/takenakasuji/go-ddd-sample/internal/domain/model/chef"
	"github.com/takenakasuji/go-ddd-sample/internal/dto"
)

type ChefRepository interface {
	List() *dto.Chefs
	ByID(id int) *dto.Chef
	Create(name string) error
	ByName(name string) *dto.Chef
}

type chefRepository struct {
	db *gorm.DB
}

func NewChefRepository(db *gorm.DB) ChefRepository {
	return &chefRepository{
		db: db,
	}
}

func (cr *chefRepository) List() *dto.Chefs {
	c := dto.NewChefs()
	cr.db.Debug().Table("chefs").Find(&c.Chefs)
	return c
}

func (cr *chefRepository) ByID(id int) *dto.Chef {
	c := dto.NewChef()
	cr.db.Debug().Table("chefs").Where("id = ?", id).Find(&c)
	return c
}

func (cr *chefRepository) Create(name string) error {
	n, _ := chef.NewName(name)
	c := chef.NewChef(n)
	cdto := c.ToDto()

	if err := cr.db.Debug().Table("chefs").Create(&cdto).Error; err != nil {
		return err
	}
	return nil
}

func (cr *chefRepository) ByName(name string) *dto.Chef {
	c := dto.NewChef()
	cr.db.Debug().Table("chefs").Where("name = ?", name).Find(&c)
	return c
}
