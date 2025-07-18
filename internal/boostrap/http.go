package boostrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-attendance/internal/config"
	"github.com/umardev500/go-attendance/internal/modules/attendance"
	"github.com/umardev500/go-attendance/internal/modules/card"
	"github.com/umardev500/go-attendance/internal/modules/device"
	"github.com/umardev500/go-attendance/internal/modules/user"
)

func ProvideFiberApp(
	config *config.Config,
	userHandler *user.Handler,
	deviceHandler *device.Handler,
	cardHandler *card.Handler,
	attendanceHandler *attendance.Handler,
) *fiber.App {
	app := fiber.New(config.Fiber)
	api := app.Group("/api")

	userHandler.Setup(api.Group("/users"))
	deviceHandler.Setup(api.Group("/devices"))
	cardHandler.Setup(api.Group("/cards"))
	attendanceHandler.Setup(api.Group("/attendances"))

	return app
}
