package boostrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-attendance/internal/config"
)

func ProvideFiberApp(config *config.Config) *fiber.App {
	app := fiber.New(config.Fiber)

	return app
}
