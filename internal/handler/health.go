package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status      string `json:"status"      example:"ok"`
	Service     string `json:"service"     example:"sbb-backend"`
	Version     string `json:"version"     example:"abc1234"`
	Environment string `json:"environment" example:"dev"`
	Timestamp   string `json:"timestamp"   example:"2026-01-01T00:00:00Z"`
}

// Health godoc
// @Summary     Health check
// @Description Returns service status, version, and environment
// @Tags        system
// @Produce     json
// @Success     200 {object} HealthResponse
// @Router      /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:      "ok",
		Service:     "sbb-backend",
		Version:     getEnv("APP_VERSION", "local"),
		Environment: getEnv("APP_ENV", "local"),
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	})
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
