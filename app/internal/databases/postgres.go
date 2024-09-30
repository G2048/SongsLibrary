package databases

import (
    "context"
    "log/slog"
    "os"
    "time"

    "SongsLibrary/app/pkg/configs"
    "SongsLibrary/app/pkg/logs"
    "github.com/jackc/pgx/v5"
)

type PostgresDatabase struct {
    dns    string
    logger *slog.Logger
}

func (p *PostgresDatabase) Connect() *pgx.Conn {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    connect, err := pgx.Connect(ctx, p.dns)
    if err != nil {
        p.logger.Error("Unable to connect to database: " + err.Error())
        os.Exit(-1)
    }

    err = connect.Ping(ctx)
    if err != nil {
        p.logger.Error("Unable to ping to database: " + err.Error())
        os.Exit(-1)
    }
    p.logger.Info("Successfully connected!")
    return connect
}
func NewPostgresDatabase() *PostgresDatabase {
    config := configs.NewPostgresConfig()
    logger := logs.NewHttpLogger(config.LogSettings)
    return &PostgresDatabase{dns: config.DSN, logger: logger.Logger}
}
