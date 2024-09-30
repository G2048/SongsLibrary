package server

import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "time"

    "SongsLibrary/app/pkg/configs"
    "SongsLibrary/app/pkg/logs"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/httplog/v2"
    httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
    Name    string
    Version string
    router  *chi.Mux
    logger  *httplog.Logger
    server  *http.Server
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func (s *Server) Start() {
    s.router.Use(middleware.Recoverer)
    s.router.Use(middleware.RequestID)
    s.router.Use(middleware.RealIP)
    s.router.Use(httplog.RequestLogger(s.logger))
    s.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    })
    s.router.Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("http://localhost:1323/swagger/doc.json"), // The url pointing to API definition
    ))

    go func() {
        err := s.server.ListenAndServe()
        switch {
        case errors.Is(err, http.ErrServerClosed):
            s.logger.Info("Stopped server")
            os.Exit(0)
        default:
            s.logger.Error(fmt.Sprintf("Failed to start server: %s", err.Error()))
            os.Exit(-1)
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
