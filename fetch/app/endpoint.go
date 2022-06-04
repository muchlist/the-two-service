package app

import (
	"fetch-api/app/handler"
	"fetch-api/app/middleware"
	"fetch-api/bussiness/repository"
	"fetch-api/bussiness/service"
	"fetch-api/conf"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func prefareEndpoint(app *fiber.App, cfg conf.Config) {

	// simple common middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept, Authorization",
	}))
	jwtMid := middleware.NewJWTMiddleware(cfg.SecretKey)

	// init repository
	cacheStore := repository.NewCurrencyStorer()
	currencyClient := repository.NewCurrencyApiCaller(cfg)
	fishClient := repository.NewFishApiCaller(cfg)

	// init service
	fishService := service.NewFetchFishServiceAssumer(fishClient, currencyClient, cacheStore)

	// init handler
	profilHandler := handler.NewProfilHandler()
	fishHandler := handler.NewFishHandler(fishService)

	// mapping url
	app.Get("/profile", jwtMid.NormalAuth(), profilHandler.DetailClaims)
	app.Get("/fish", jwtMid.NormalAuth(), fishHandler.FindFish)
}
