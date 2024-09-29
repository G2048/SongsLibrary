package main

import (
    "SongsLibrary/src/internal/databases/migrations"
    "SongsLibrary/src/pkg/configs"
    "SongsLibrary/src/server"
    v1 "SongsLibrary/src/server/routers/v1"
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
