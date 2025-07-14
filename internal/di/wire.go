//go:build wireinject
// +build wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/umardev500/go-attendance/internal/boostrap"
	"github.com/umardev500/go-attendance/internal/config"
)

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var AppSet = wire.NewSet(
	boostrap.ProvideFiberApp,
)

func InitializeFiberApp(config *config.Config) *fiber.App {
	wire.Build(AppSet)
	return nil
}
