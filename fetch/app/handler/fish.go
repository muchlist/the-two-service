package handler

import (
	"fetch-api/bussiness/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewFishHandler(service service.FetchFishServiceAssumer) *FishHandler {
	return &FishHandler{
		Service: service,
	}
}

type FishHandler struct {
	Service service.FetchFishServiceAssumer
}

func (f *FishHandler) FindFish(c *fiber.Ctx) error {

	result, err := f.Service.FetchData()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  result,
		"error": nil,
	})
}

func (f *FishHandler) FindFishAggregate(c *fiber.Ctx) error {

	result, err := f.Service.GetAggregatedData()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  result,
		"error": nil,
	})
}
