package storage

import (
	"github.com/golanguzb71/microservice-project-template/config"
	"github.com/golanguzb71/microservice-project-template/pkg/db"
	"github.com/golanguzb71/microservice-project-template/pkg/logger"
	"github.com/golanguzb71/microservice-project-template/storage/postgres"
	"github.com/golanguzb71/microservice-project-template/storage/repo"
)

type StorageI interface {
	HealthCheck() repo.HealthCheckRepo
}

type storagePg struct {
	healthCheck repo.HealthCheckRepo
}

func New(db *db.Postgres, log logger.Logger, cfg *config.Config) StorageI {
	return &storagePg{
		healthCheck: postgres.NewHealthCheckRepo(db, log, cfg),
	}
}

func (s *storagePg) HealthCheck() repo.HealthCheckRepo {
	return s.healthCheck
}
