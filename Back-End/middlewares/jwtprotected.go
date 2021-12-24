package middlewares

import (
	"back-end/config"
	"github.com/gofiber/fiber/v2"
	jwtmiddleware "github.com/gofiber/jwt/v2"


)

// JWTProtected middleware
func JWTProtected() func(*fiber.Ctx) error {
	return jwtmiddleware.New(jwtmiddleware.Config{
		SigningKey:   []byte(config.GetKey("APP_SECRET")),
		//ContextKey: "key",
	})
}
