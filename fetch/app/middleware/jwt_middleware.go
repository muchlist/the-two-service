package middleware

import (
	"errors"
	"fetch-api/pkg/mjwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	headerKey = "Authorization"
	bearerKey = "Bearer"
)

type JWTMiddleware struct {
	Secret string
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{
		Secret: secret,
	}
}

// NormalAuth memerlukan salah satu role yang tertulis agar diloloskan ke proses berikutnya
func (j *JWTMiddleware) NormalAuth(rolesReq ...string) fiber.Handler {

	jwtReader := mjwt.New(j.Secret)

	return func(c *fiber.Ctx) error {
		authHeader := c.Get(headerKey)
		claims, err := authRoleValidator(jwtReader, authHeader, rolesReq)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error(), "data": nil})
		}

		c.Locals(mjwt.CLAIMS, claims)
		return c.Next()
	}
}

func authRoleValidator(jwtReader mjwt.TokenReader, authHeader string, rolesRequired []string) (mjwt.CustomClaim, error) {

	if !strings.Contains(authHeader, bearerKey) {
		return mjwt.CustomClaim{}, errors.New("Unauthorized")
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		return mjwt.CustomClaim{}, errors.New("Unauthorized")
	}

	token, err := jwtReader.ValidateToken(tokenString[1])
	if err != nil {
		return mjwt.CustomClaim{}, err
	}

	claims, err := jwtReader.ReadToken(token)
	if err != nil {
		return mjwt.CustomClaim{}, err
	}

	// if required role is 0 , access granted
	if len(rolesRequired) == 0 {
		return claims, nil
	}

	// if required role is available, do loop for check claims role
	if len(rolesRequired) != 0 {
		for _, roleReq := range rolesRequired {
			if strings.EqualFold(roleReq, claims.Role) {
				return claims, nil
			}
		}
	}

	return mjwt.CustomClaim{}, fmt.Errorf("unauthorized, role %s needed", rolesRequired)
}
