package main

import (
	"github.com/ak-yudha/lili_test/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/product", handlers.CreateProduct)
	app.Get("/product", handlers.ListProduct)
	app.Post("/user", handlers.CreateUser)
	app.Post("/favorites", handlers.CreateFavorites)
}
