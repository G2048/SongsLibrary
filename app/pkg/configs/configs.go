package configs

import (
    "os"

    "github.com/joho/godotenv"
)

type AppSettings struct {
    AppName    string
    AppVersion string
}

type LogSettings struct {
    AppSettings
    LogLevel string `short:"l" help:"log level" default:"debug"`
}

type ServerSettings struct {
    LogSettings
    Port string `short:"p" help:"Server port" default:"3333"`
    Host string `short:"h" help:"Server host" default:"127.0.0.1"`
}

func (s ServerSettings) FullServerAddress() string {
    return s.Host + ":" + s.Port
}
func NewServerSettings() *ServerSettings {
    env, err := godotenv.Read()
    if err != nil {
        return &ServerSettings{
            LogSettings: LogSettings{
                AppSettings: AppSettings{
                    AppName:    os.Getenv("appname"),
                    AppVersion: os.Getenv("appversion"),
                },
                LogLevel: os.Getenv("logLevel"),
            },
            Port: os.Getenv("port"),
            Host: os.Getenv("host"),
        }
    }

    return &ServerSettings{
        LogSettings: LogSettings{
            AppSettings: AppSettings{
                AppName:    env["appname"],
                AppVersion: env["appversion"],
            },
            LogLevel: env["logLevel"],
        },
        Port: env["port"],
        Host: env["host"],
    }
}

type PostgresConfig struct {
    LogSettings
    DSN string
}

func NewPostgresConfig() *PostgresConfig {
    env, err := godotenv.Read()

    if err != nil {
        // panic(err)
        return &PostgresConfig{
            LogSettings: LogSettings{
                AppSettings: AppSettings{
                    AppName:    os.Getenv("appname"),
                    AppVersion: os.Getenv("appversion"),
                },
                LogLevel: os.Getenv("logLevel"),
            },
            DSN: os.Getenv("dsn"),
        }
    }

    return &PostgresConfig{
        LogSettings: LogSettings{
            AppSettings: AppSettings{
                AppName:    env["appname"],
                AppVersion: env["appversion"],
            },
            LogLevel: env["logLevel"],
        },
        DSN: env["dsn"],
    }
}
