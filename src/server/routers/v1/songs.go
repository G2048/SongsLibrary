package v1

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "strings"

    "SongsLibrary/src/server"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
)

type ErrResponse struct {
    Err            error `json:"-"` // low-level runtime error
    HTTPStatusCode int   `json:"-"` // http response status code

    StatusText string `json:"status"`          // user-level status message
    AppCode    int64  `json:"code,omitempty"`  // application-specific error code
    ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, e.HTTPStatusCode)
    return nil
}
func ErrInvalidRequest(err error) render.Renderer {
    return &ErrResponse{
        Err:            err,
        HTTPStatusCode: 400,
        StatusText:     "Invalid request.",
        ErrorText:      err.Error(),
    }
}

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

    r.Get("/", get)
    r.Post("/", create)
    r.Get("/{id}", getId)
    r.Patch("/{id}", update)
    r.Delete("/{id}", delete)

    return &SongRouters{
        version: "v1",
        name:    "songs",
        router:  r,
    }
}

type SongsRequest struct {
    Title  string `json:"title"`
    Artist string `json:"artist",omitempty`
}

func (s *SongsRequest) Bind(r *http.Request) error {
    if s.Artist == "" {
        return errors.New("missing required Artist fields.")
    }
    s.Title = strings.ToLower(s.Title)
    return nil
}

func (s *SongsRequest) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, http.StatusCreated)
    err := json.NewDecoder(r.Body).Decode(s)
    if err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
    }
    // w.Write([]byte(fmt.Sprintf("song: %s\n", s.Title)))
    // w.Write([]byte(fmt.Sprintf("artist: %s\n", s.Artist)))

    return nil
}

func (s *SongsRequest) RenderResponse() render.Renderer {
    return s
}

func get(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("song 1"))
}
func getId(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(fmt.Sprintf("song %s", chi.URLParam(r, "id"))))
}
func create(w http.ResponseWriter, r *http.Request) {
    var data SongsRequest
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
    }
    w.Write([]byte(fmt.Sprintf("song: %s\n", data.Title)))
    w.Write([]byte(fmt.Sprintf("artist: %s\n", data.Artist)))
}
func update(w http.ResponseWriter, r *http.Request) {
    var data SongsRequest
    r.Body = http.MaxBytesReader(w, r.Body, 1024)
    if err := render.Bind(r, &data); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }
    render.Status(r, http.StatusCreated)
    render.Render(w, r, &data)
}
func delete(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(fmt.Sprintf("song %s", chi.URLParam(r, "id"))))
}
