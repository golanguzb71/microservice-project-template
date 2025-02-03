package postgres

import (
	"context"
	"github.com/golanguzb71/microservice-project-template/config"
	pb "github.com/golanguzb71/microservice-project-template/genproto/template_service"
	"github.com/golanguzb71/microservice-project-template/pkg/db"
	"github.com/golanguzb71/microservice-project-template/pkg/logger"
	"github.com/golanguzb71/microservice-project-template/storage/repo"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type healthcheck struct {
	db  *db.Postgres
	cff *config.Config
	log logger.Logger
}

func NewHealthCheckRepo(db *db.Postgres, log logger.Logger, cff *config.Config) repo.HealthCheckRepo {
	return &healthcheck{db: db, log: log, cff: cff}
}

func (h *healthcheck) HealthCheck() (*pb.HealthCheckResponse, error) {
	id := uuid.New().String()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	tr, err := h.db.Db.Begin(ctx)
	if err != nil {
		h.log.Error("HealthCheck: " + err.Error())
		return nil, err
	}

	queryBuild := h.db.Builder.Insert("healthcheck").Columns("id").Values(id)
	query, args, err := queryBuild.ToSql()
	if err != nil {
		tr.Rollback(ctx)
		h.log.Error("HealthCheck: " + err.Error())
		return nil, err
	}

	_, err = tr.Exec(ctx, query, args...)
	if err != nil {
		tr.Rollback(ctx)
		h.log.Error("HealthCheck: " + err.Error())
		return nil, err
	}

	queryDelete := h.db.Builder.Delete("healthcheck").Where(squirrel.Eq{"id": id})
	query, args, err = queryDelete.ToSql()
	if err != nil {
		h.log.Error("HealthCheck: " + err.Error())
		return nil, err
	}

	_, err = tr.Exec(ctx, query, args...)
	if err != nil {
		tr.Rollback(ctx)
		h.log.Error("HealthCheck: " + err.Error())
		return nil, err
	}

	err = tr.Commit(ctx)
	if err != nil {
		h.log.Error("HealthCheck: " + err.Error())
		return nil, err
	}

	return &pb.HealthCheckResponse{
		Healthy:       true,
		UnHealthLevel: "",
	}, nil
}
