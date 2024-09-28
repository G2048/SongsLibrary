package main

import (
    "SongsLibrary/src/server"
    v1 "SongsLibrary/src/server/routers/v1"
)
import "SongsLibrary/src/pkg/configs"

func main() {
    settings := *configs.NewServerSettings()

    s := server.NewServer(settings)
    s.Start()
    s.AddRouter(v1.NewSongRouters())

    // Blocking operation
    s.Stop()
}
