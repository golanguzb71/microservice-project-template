package repo

import pb "github.com/golanguzb71/microservice-project-template/genproto/template_service"

type HealthCheckRepo interface {
	HealthCheck() (*pb.HealthCheckResponse, error)
}
