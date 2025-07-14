package boostrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-attendance/internal/config"
	"github.com/umardev500/go-attendance/internal/modules/user"
)

func ProvideFiberApp(
	config *config.Config,
	userHandler *user.Handler,
) *fiber.App {
	app := fiber.New(config.Fiber)
	api := app.Group("/api")
	userHandler.Setup(api.Group("/users"))

	return app
}
