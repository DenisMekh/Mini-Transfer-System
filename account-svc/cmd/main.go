package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/config"
	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/handler"
	rp "github.com/DenisMekh/mini-transfer-system/account-svc/internal/repo/postgres"
	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/service"
	"github.com/DenisMekh/mini-transfer-system/account-svc/pkg/logger"
	"github.com/DenisMekh/mini-transfer-system/account-svc/pkg/postgres"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New(cfg)
	ctx := context.Background()
	pool, err := postgres.New(ctx, &cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database")
	}
	rep := rp.NewAccountRepo(pool)
	srv := service.NewAccountService(rep)
	han := handler.NewAccountHandler(srv)
	s, err := New(han, fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatal("Failed to start server")
	}
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := s.Run(); err != nil {
			logger.Fatal("Failed to start server")
		}
	}()
	<-ctx.Done()
	logger.Info("Shutting down")
	s.Stop()
}
