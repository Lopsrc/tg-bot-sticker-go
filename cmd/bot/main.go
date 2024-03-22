package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"tg-bot-sticker-go/internal/bot"
	"tg-bot-sticker-go/internal/config"
	stickerhandle "tg-bot-sticker-go/internal/handlers/sticker"
	"tg-bot-sticker-go/internal/storage/redis"
)

const (
	pathConfig 			= "config/local.yaml"
	envLocal   			= "local"
	envDev     			= "dev"
	envProd    			= "prod"
)

func main() {
	// Create a config.
	cfg := config.MustLoadPath(pathConfig)
	// Create a logger.
	log := setupLogger(cfg.Env)
	// Create a bot.
	b := bot.New(cfg.PathEnv)
	// Create a redis client.
	client := redis.New(cfg)
	// Handle.
	stickerHandle := stickerhandle.New(b, client, log)
	stickerHandle.Register()
	// Start the bot.
	go func() {
		b.Start()
	}()
	log.Info("Bot started.")	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	b.Stop()
	log.Info("Bot stopped.")	
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}