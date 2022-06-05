package handler

import (
	"fetch-api/pkg/mjwt"

	"github.com/gofiber/fiber/v2"
)

func NewProfilHandler() *ProfilHandler {
	return &ProfilHandler{}
}

type ProfilHandler struct{}

func (p *ProfilHandler) DetailClaims(c *fiber.Ctx) error {
	claims, _ := c.Locals(mjwt.CLAIMS).(mjwt.CustomClaim)
	return c.JSON(fiber.Map{
		"data":  claims,
		"error": nil,
	})
}
