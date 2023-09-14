package main

import (
	_ "github.com/ak-yudha/lili_test/docs"
	"github.com/ak-yudha/lili_test/handlers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/product", handlers.CreateProduct)
	app.Get("/product", handlers.ListProduct)
	app.Post("/user", handlers.CreateUser)
	app.Post("/favorites", handlers.CreateFavorites)
	app.Get("/product/:id", handlers.GetProduct)

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}
