package main

import (
	"context"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-attendance/internal/config"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/di"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	config := config.NewConfig()
	app := di.InitializeFiberApp(config)

	// Auto migration
	client := database.NewEntClient(config)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	log.Info().Msgf("Listening on port %d", config.Port)

	app.Listen(":" + strconv.Itoa(config.Port))
}
