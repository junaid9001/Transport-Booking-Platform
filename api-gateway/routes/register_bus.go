package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junaid9001/tripneo/api-gateway/config"
	"github.com/junaid9001/tripneo/api-gateway/proxy"
	"github.com/redis/go-redis/v9"
)

func RegisterBusRoutes(app *gin.Engine, cfg *config.Config, rdb *redis.Client) {
	api := app.Group("/api/bus")

	api.GET("/health", proxy.To(cfg.BUS_SERVICE_URL))
}
