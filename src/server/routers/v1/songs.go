package v1

import (
    "fmt"
    "net/http"

    "SongsLibrary/src/server"
    "github.com/go-chi/chi/v5"
)

type SongRouters struct {
    version string
    name    string
    router  *chi.Mux
}

func (s *SongRouters) Path() string {
    return "/" + s.version + "/" + s.name
}
func (s *SongRouters) Router() *chi.Mux {
    return s.router
}

func NewSongRouters() server.RoutersInterface {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("song 1"))
    })

    r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(fmt.Sprintf("song %s", chi.URLParam(r, "id"))))
    })

    return &SongRouters{
        version: "v1",
        name:    "songs",
        router:  r,
    }
}
