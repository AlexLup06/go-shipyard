package main

import (
	"__MODULE__/internal/bootstrap"
	"__MODULE__/internal/config"
	httpserver "__MODULE__/internal/http"
	"__MODULE__/internal/http/kit/render"
	"__MODULE__/internal/http/middleware"
	"__MODULE__/internal/logging"
	"__MODULE__/internal/store"
	"__MODULE__/internal/store/schema"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Version = "dev"

func main() {
	// Binary self-check for Docker HEALTHCHECK
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		os.Exit(0)
	}

	// regular server
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, err := logging.New(cfg.Logging.Level)
	if err != nil {
		logger.Error("invalid token key configuration", "err", err)
		os.Exit(1)
	}
	logger.Info("starting __APP_SLUG__")

	store, err := store.New(store.Config{
		Host:            cfg.DB.Host,
		Port:            cfg.DB.Port,
		Username:        cfg.DB.Username,
		Password:        cfg.DB.Password,
		Database:        cfg.DB.Database,
		Schema:          cfg.DB.Schema,
		Timezone:        cfg.DB.Timezone,
		LogSql:          cfg.DB.LogSQL,
		MaxOpenConns:    cfg.DB.MaxOpenConns,
		MaxIdleConns:    cfg.DB.MaxIdleConns,
		ConnMaxLifetime: cfg.DB.ConnMaxLifetime,
		ConnMaxIdleTime: cfg.DB.ConnMaxIdleTime,
	})
	if err != nil {
		logger.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	err = bootstrap.CheckSchemaVersion(store, schema.RequiredSchemaVersion)
	if err != nil {
		logger.Error("startup failed", "err", err)
		os.Exit(1)
	}

	// txManager := tx.New(store)

	mw := middleware.Middlewares{}

	assets, err := render.LoadAssetsManifest("./internal/http/static/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	renderer := render.New(assets)

	server := httpserver.NewServer(httpserver.ServerConfig{
		Version: Version,
		Addr:    cfg.Values.HttpAddr,
		Dev:     cfg.Values.AppEnv == "dev",
		Logger:  logger,
		Store:   store,
		Render:  renderer,
	}, mw)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("http server listening", "addr", cfg.Values.HttpAddr)
		if err := server.Start(); err != nil {
			logger.Error("http server stopped unexpectedly", "err", err)
			stop()
		}
	}()

	<-ctx.Done()

	logger.Info("shutting down __APP_SLUG__")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "err", err)
	}

	logger.Info("__APP_SLUG__ stopped")
}
