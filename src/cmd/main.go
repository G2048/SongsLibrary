package main

import "SongsLibrary/src/server"
import "SongsLibrary/src/pkg/configs"

func main() {
    settings := *configs.NewServerSettings()
    s := server.NewServer(settings)
    s.Start()

    // Blocking operation
    s.Stop()
}
