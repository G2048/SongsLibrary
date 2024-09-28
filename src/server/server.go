package server

import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "time"

    "SongsLibrary/src/pkg/configs"
    "SongsLibrary/src/pkg/logs"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/httplog/v2"
)

type Server struct {
    Name    string
    Version string
    router  *chi.Mux
    logger  *httplog.Logger
    server  *http.Server
}

func (s *Server) Start() {

    s.router.Use(middleware.Recoverer)
    s.router.Use(middleware.RequestID)
    s.router.Use(middleware.RealIP)
    s.router.Use(httplog.RequestLogger(s.logger))
    s.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    })

    go func() {
        err := s.server.ListenAndServe()
        switch {
        case errors.Is(err, http.ErrServerClosed):
            s.logger.Info("Stopped server")
            os.Exit(0)
        default:
            s.logger.Error(fmt.Sprintf("Failed to start server: %s", err.Error()))
        }
    }()
}

type RoutersInterface interface {
    Path() string
    Router() *chi.Mux
}

func (s *Server) AddRouter(router RoutersInterface) {
    s.router.Mount(router.Path(), router.Router())
}

func (s *Server) Stop() {
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, os.Interrupt, os.Kill)
    <-sigc

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

    err := s.server.Shutdown(ctx)
    if err != nil {
        s.logger.Error(fmt.Sprintf("Failed to shutdown server: %s", err.Error()))
        os.Exit(1)
    }
    s.logger.Info("Stopped server")
    os.Exit(0)
}

func NewServer(settings configs.ServerSettings) *Server {
    logger := logs.NewHttpLogger(settings.LogSettings)
    logger.Info("Starting server")
    logger.Debug(fmt.Sprintf("Server Port: %s", settings.Port))

    router := chi.NewRouter()
    server := &http.Server{
        Addr:    settings.FullServerAddress(),
        Handler: router,
    }

    return &Server{
        router: router,
        logger: logger,
        server: server,
    }
}

// func RequestLogger(logger *httplog.Logger, strings []string) http.Handler {
//
// }
