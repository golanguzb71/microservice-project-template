package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/golanguzb71/microservice-project-template/genproto/template_service"
)

func (h *Handler) HealthCheck(ctx *gin.Context) {
	res, err := h.service.HealthCheckService.HealthCheck(ctx, &pb.Empty{})
	if err != nil {
		ctx.JSON(500, res)
		return
	}

	if !res.Healthy {
		ctx.JSON(500, res)
		return
	}

	ctx.JSON(200, res)
}
