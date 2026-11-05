package main

import (
	"fmt"
	"net/http"

	"github.com/3115826227/Anduin/internal/config"
	"github.com/3115826227/Anduin/internal/database"
	"github.com/3115826227/Anduin/internal/logger"
	"github.com/3115826227/Anduin/internal/server"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	log, err := logger.New(cfg.Log.Level, cfg.Log.Format)
	if err != nil {
		return err
	}
	defer log.Sync()

	if cfg.Database.Enabled {
		db, err := database.Open(cfg.Database.DSN)
		if err != nil {
			return err
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		defer sqlDB.Close()
		log.Info("database connected")
	}

	router := server.New(cfg.Server.Mode)
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Info("Anduin starting", zap.String("address", address))
	return http.ListenAndServe(address, router)
}
