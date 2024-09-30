package main

import (
    "SongsLibrary/app/internal/databases/migrations"
    "SongsLibrary/app/pkg/configs"
    "SongsLibrary/app/server"
    "SongsLibrary/app/server/routers/v1"
)

func main() {
    settings := *configs.NewServerSettings()

    migrations.RunMigrations()

    // Start Server
    s := server.NewServer(settings)
    s.Start()
    s.AddRouter(v1.NewSongRouters())

    // Blocking operation
    s.Stop()
}
