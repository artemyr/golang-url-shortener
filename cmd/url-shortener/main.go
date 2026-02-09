package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	mwLogger "url-shortener/internal/http-server/middleware/logger"
	"url-shortener/internal/lib/logger/handlers/slogpretty"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	//     log.With(slog.String("env", cfg.Env))
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// url, err := storage.GetURL("google")
	// if err != nil {
	// 	log.Error("failed to get url", sl.Err(err))
	// 	os.Exit(1)
	// }

	// log.Info("saved url is", slog.String("url", url))

	// url, err = storage.GetURL("error")
	// if err != nil {
	// 	log.Error("failed to get url", sl.Err(err))
	// 	os.Exit(1)
	// }

	// id, err := storage.SaveURL("https://google.com", "google")
	// if err != nil {
	//     log.Error("failed to save url", sl.Err(err))
	//     os.Exit(1)
	// }

	// log.Info("saved url", slog.Int64("id", id))

	// id, err = storage.SaveURL("https://google.com", "google")
	// if err != nil {
	//     log.Error("failed to save url", sl.Err(err))
	//     os.Exit(1)
	// }

	_ = storage
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
