package main

import (
	"github.com/golanguzb71/microservice-project-template/config"
	"github.com/golanguzb71/microservice-project-template/pkg/logger"
	"github.com/golanguzb71/microservice-project-template/server/grpc"
	httpserver "github.com/golanguzb71/microservice-project-template/server/http"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.Environment, "name space for monitoring service log")

	grpcServices, err := grpc.New(cfg, log)
	if err != nil {
		log.Error("Error while initializing grpcServices", logger.Error(err))
		return
	}

	httpServer, err := httpserver.New(cfg, log, grpcServices)
	if err != nil {
		log.Error("Error while initializing http server", logger.Error(err))
		return
	}

	go func() {
		err := httpServer.Run(log, cfg)
		if err != nil {
			log.Fatal("Error while running http server", logger.Error(err))
		}
	}()

	grpcServices.Run(log, cfg)
}
