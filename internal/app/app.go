package app

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Slobodskoy/calc_go/internal/api"
	"github.com/Slobodskoy/calc_go/internal/config"
	"github.com/Slobodskoy/calc_go/internal/middleware"
)

type App struct {
	cfg config.Config
}

func New(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Run() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	handler := new(api.CalcHandler)
	handlerWithLog := middleware.AccessLog(handler.Calc)
	handlerWithRecovery := middleware.Recovery(handlerWithLog)
	http.HandleFunc("POST /", handlerWithRecovery)
	slog.Info("start http server", slog.Int("port", a.cfg.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.Port), nil); err != nil {
		log.Fatal(err)
	}
}
