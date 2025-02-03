package services

import (
	"github.com/golanguzb71/microservice-project-template/config"
	"github.com/golanguzb71/microservice-project-template/pkg/logger"
	"github.com/golanguzb71/microservice-project-template/server/grpc/client"
	"github.com/golanguzb71/microservice-project-template/storage"
)

type ServiceOptions struct {
	ServiceManager client.ServiceManager
	Storage        storage.StorageI
	Config         *config.Config
	Logger         logger.Logger
}
