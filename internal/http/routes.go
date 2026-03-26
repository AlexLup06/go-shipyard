package http

import (
	"net/http"
	"time"

	"__MODULE__/internal/http/handlers/meta"
	httpmiddleware "__MODULE__/internal/http/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg ServerConfig, mw httpmiddleware.Middlewares) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(httpmiddleware.RequestLogger(cfg.Logger))

	registerRoutes(r, cfg, mw)

	return r
}

func registerRoutes(r chi.Router, cfg ServerConfig, mw httpmiddleware.Middlewares) {
	r.Group(func(r chi.Router) {
		r.Get("/health", meta.Health)
		r.Get("/version", meta.Version(cfg.Version))
	})

	meta.RegisterStatic(r, meta.StaticConfig{Dev: cfg.Dev})
}
