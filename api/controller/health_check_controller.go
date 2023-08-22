package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type HealthCheckControllerImpl struct {
}

type HealthCheckRes struct {
	Message string `json:"message" example:"pong"`
}

// @BasePath /api/v1

// HealthCheck godoc
// @Summary Health check
// @Schemes
// @Description Ping
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} HealthCheckRes
// @Router /ping [get]
func (h HealthCheckControllerImpl) HealthCheck(c *gin.Context) {
	resJson := HealthCheckRes{
		Message: "pong",
	}
	c.JSON(http.StatusOK, resJson)
}

func HealthCheckControllerInit() *HealthCheckControllerImpl {
	return &HealthCheckControllerImpl{}
}
