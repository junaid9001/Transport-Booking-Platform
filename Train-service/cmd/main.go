package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/nabeel-mp/tripneo/train-service/config"
	"github.com/nabeel-mp/tripneo/train-service/db"
	"github.com/nabeel-mp/tripneo/train-service/redis"
	"github.com/nabeel-mp/tripneo/train-service/routes"
)

func main() {
	// 1. Load config from .env
	cfg := config.LoadConfig()

	// 2. Connect to PostgreSQL
	db.ConnectPostgres(cfg)

	// 3. Connect to Redis
	rdb := redis.Client(cfg.REDIS_HOST, cfg.REDIS_PORT)

	// 4. Kafka producer — initialized in Phase 6
	// kafka.InitProducer(cfg)

	// 5. gRPC server — started in Phase 5
	// go grpc.StartServer(cfg)

	// 6. Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "TripNEO Train Service v1.0",
	})

	// 7. Register all routes
	routes.Register(app, cfg, rdb)

	// 8. Start HTTP server
	log.Printf("Train service starting on port %s", cfg.APP_PORT)
	if err := app.Listen(":" + cfg.APP_PORT); err != nil {
		log.Fatal(err)
	}
}
