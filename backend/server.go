package main

import (
	controller "go-todo/backend/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/zadanie/:id", controller.Zadanie)
	e.GET("/zadania", controller.Zadania)
	e.POST("/zadanie/dodaj", controller.Dodaj)
	e.DELETE("/zadanie/usun/:id", controller.Usun)
	e.PUT("/zadanie/zakoncz/:id", controller.Zakoncz)

	e.Logger.Fatal(e.Start(":2381"))
}
