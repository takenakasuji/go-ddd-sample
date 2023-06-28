package main

import (
	"github.com/labstack/echo/v4"
	"github.com/takenakasuji/go-ddd-sample/internal/application"
	"github.com/takenakasuji/go-ddd-sample/internal/handler"
	"github.com/takenakasuji/go-ddd-sample/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := "docker:docker@tcp(127.0.0.1:3307)/zidane?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	chefRepository := repository.NewChefRepository(db)
	e.GET("/chefs", handler.Chefs(application.NewChefApplicationService(chefRepository)))
	e.GET("/chefs/:id", handler.Chef(application.NewChefApplicationService(chefRepository)))
	e.POST("/chefs", handler.CreateChef(application.NewChefApplicationService(chefRepository)))

	if err := e.Start(":8888"); err != nil {
		panic("echo does not work.")
	}
}
