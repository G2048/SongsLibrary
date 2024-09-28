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
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.RequestID)
    s.router.Use(middleware.RealIP)
    s.router.Use(httplog.RequestLogger(s.logger))
    r.Use(middleware.Heartbeat("/ping"))

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
func (s *Server) AddRouter() {
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

    server := &http.Server{
        Addr: settings.FullServerAddress(),
    }

    return &Server{
        router: chi.NewRouter(),
        logger: logger,
        server: server,
    }
}
