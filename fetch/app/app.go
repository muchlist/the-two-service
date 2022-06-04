package app

import (
	"fetch-api/conf"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func RunApp() {

	// load config
	cfg := conf.Load()

	// create fiber app
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
				// Send custom error page
				ctx.Status(code).JSON(fiber.Map{
					"data":  nil,
					"error": err.Error(),
				})
				return nil
			},
		},
	)

	// gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// do fullfill dependency and routing
	prefareEndpoint(app, *cfg)

	// blocking and listen for fiber app
	if err := app.Listen(":8081"); err != nil {
		log.Panic()
	}
	// cleanup app
	fmt.Println("Running cleanup tasks...")
}
