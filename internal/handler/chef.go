package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/takenakasuji/go-ddd-sample/internal/application"
)

func Chefs(chefApp application.ChefApplicationService) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, chefApp.List())
	}
}

func Chef(chefApp application.ChefApplicationService) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		sid := c.Param("id")
		return c.JSON(http.StatusOK, chefApp.ChefProfile(sid))
	}
}

func CreateChef(chefApp application.ChefApplicationService) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		err := chefApp.CreateChef(name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to new chef")
		}
		return c.JSON(http.StatusOK, "the resource is created")
	}
}
