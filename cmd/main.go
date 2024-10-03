package main

import (
	"ITForce/api"
	_ "ITForce/docs"
	"ITForce/internal/config"
	"ITForce/internal/repositories"
	"ITForce/internal/repositories/migrations"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

// @title itForce
// @version 1.0
// @description Тестовое задание.

// @contact.name Константин
// @contact.url https://t.me/London68
func main() {

	//Инициализация
	fmt.Println("Чтение конфигурации")
	cfg := config.New()
	fmt.Println("Подключение к базе данных")
	db := repositories.New(cfg.DB)
	_, err := migrations.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.POST("/purchase", func(c echo.Context) error {
		return api.PurchaseHandler(c, db)
	})
	e.GET("/history/:id", func(c echo.Context) error {
		return api.HistoryHandler(c, db)
	})
	e.GET("/history", func(c echo.Context) error {
		return api.HistoryHandler(c, db)
	})
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.API.Port)))
}
