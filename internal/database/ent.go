package database

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-attendance/internal/config"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/internal/ent/migrate"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to postgres")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func NewEntClient(config *config.Config) *ent.Client {
	ctx := context.Background()
	db := config.DB

	dsn := db.DSN()

	client := Open(dsn)
	if err := client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	log.Info().
		Str("dsn", dsn).
		Msg("database connected successfully")

	return client
}
