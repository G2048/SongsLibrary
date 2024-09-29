package migrations

import (
    "database/sql"
    "embed"
    "errors"
    "fmt"
    "log"

    "SongsLibrary/src/pkg/configs"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    "github.com/golang-migrate/migrate/v4/source"
    "github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migrator struct {
    srcDriver source.Driver
}

func NewMigrator(sqlFiles embed.FS, dirName string) *Migrator {
    driver, err := iofs.New(sqlFiles, dirName)
    if err != nil {
        panic(err)
    }
    return &Migrator{
        srcDriver: driver,
    }
}

func (m *Migrator) ApplyMigrations(db *sql.DB) error {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("unable to create db instance: %v", err)
    }

    migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "psql_db", driver)
    if err != nil {
        return fmt.Errorf("unable to create migration: %v", err)
    }

    defer func() {
        srcErr, dbErr := migrator.Close()
        if srcErr != nil {
            log.Println(srcErr)
        }
        if dbErr != nil {
            log.Println(dbErr)
        }
    }()

    if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
        return fmt.Errorf("unable to apply migrations %v", err)
    }

    return nil
}

const migrationsDir = "."

//go:embed *.sql
var MigrationsFS embed.FS

func RunMigrations() {
    settings := configs.NewPostgresConfig()
    // Run Migrations
    migrator := NewMigrator(MigrationsFS, migrationsDir)
    log.Printf("Running migrations for %s", settings.DSN)

    conn, err := sql.Open("postgres", settings.DSN)
    if err != nil {
        panic(err)
    }

    defer func() {
        err := conn.Close()
        if err != nil {
            panic(err)
        }
    }()

    // Apply Migrations
    err = migrator.ApplyMigrations(conn)
    if err != nil {
        panic(err)
    }
    log.Println("Migrations applied!!")
}
