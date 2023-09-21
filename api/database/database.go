package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const MIGRATION_DIRECTORY = "postgres/migration"

type PgxPool interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Close()
}

type PostgresDatabase struct {
	Pool             PgxPool
	connectionString string
}

func NewPostgresDatabase() (*PostgresDatabase, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal().Msgf("unable to load .env file: %v", err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", // pragma: allowlist secret
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL_MODE"),
	)
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &PostgresDatabase{
		Pool:             pool,
		connectionString: connectionString,
	}, nil
}

func (postgres *PostgresDatabase) Migrate() error {
	projectRootPath, _ := os.Getwd()
	migrationsPath := filepath.Join(projectRootPath, MIGRATION_DIRECTORY)
	migrationSourceUrl := fmt.Sprintf("file://%v", migrationsPath)
	migration, err := migrate.New(
		migrationSourceUrl,
		postgres.connectionString,
	)
	if err != nil {
		return fmt.Errorf("unable to create migrate instance: %v", err)
	}

	err = migration.Up()
	if err != nil && err.Error() != "no change" {
		return fmt.Errorf("unable to apply migrations to database: %v", err)
	}

	return nil
}

func (postgres *PostgresDatabase) Close() {
	postgres.Pool.Close()
}
