package app

import (
	"fetch-api/app/handler"
	"fetch-api/bussiness/middleware"
	"fetch-api/conf"

	"github.com/gofiber/fiber/v2"
)

func prefareEndpoint(app *fiber.App, cfg conf.Config) {

	// custom middleware
	jwtMid := middleware.NewJWTMiddleware(cfg.SecretKey)

	// init handler
	profilHandler := handler.NewProfilHandler()

	// mapping url
	app.Get("/profile", jwtMid.NormalAuth(), profilHandler.DetailClaims)
}
