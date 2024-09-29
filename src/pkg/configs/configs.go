package configs

import "github.com/joho/godotenv"

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
        panic(err)
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
