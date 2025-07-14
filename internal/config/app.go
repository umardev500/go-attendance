package config

import "github.com/gofiber/fiber/v2"

type Config struct {
	Port  int
	DB    DBConfig
	Fiber fiber.Config
}

func NewConfig() *Config {
	dbConfig := NewDBConfig()

	return &Config{
		Port: 3000,
		DB:   dbConfig,
		Fiber: fiber.Config{
			DisableStartupMessage: true,
		},
	}
}
