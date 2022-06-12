package app

import (
	"fetch-api/conf"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

const version = "1.0.0"

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

	// start debug server
	debugPort := "4000"
	if cfg.DebugPort != "" {
		debugPort = cfg.DebugPort
	}
	debugMux := debugMux()
	go func(mux *http.ServeMux) {
		if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", debugPort), mux); err != nil {
			log.Print("serve debug api", err)
		}
	}(debugMux)

	// do fullfill dependency and routing
	prefareEndpoint(app, *cfg)

	// gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// blocking and listen for fiber app
	port := "8081"
	if cfg.ServerPort != "" {
		port = cfg.ServerPort
	}
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Panic()
	}
	// cleanup app
	fmt.Println("Running cleanup tasks...")
}
