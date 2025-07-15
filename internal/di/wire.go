//go:build wireinject
// +build wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/umardev500/go-attendance/internal/boostrap"
	"github.com/umardev500/go-attendance/internal/config"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/modules/card"
	"github.com/umardev500/go-attendance/internal/modules/device"
	"github.com/umardev500/go-attendance/internal/modules/user"
)

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var AppSet = wire.NewSet(
	boostrap.ProvideFiberApp,
	user.Set,
	device.Set,
	card.Set,
	ProvideValidator,
	database.NewTransactionManager,
	database.NewEntClient,
)

func InitializeFiberApp(config *config.Config) *fiber.App {
	wire.Build(AppSet)
	return nil
}
