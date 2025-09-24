package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/database"
	"github.com/arthu/shop-api-go/api/handler"
)

func main() {
	// Load env
	_ = os.Setenv("GIN_MODE", "release")
	cfg := config.Load()

	// Connect DB
	if err := database.Connect(cfg.DBDSN); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer database.Close()

	r := gin.Default()
	handler.RegisterRoutes(r, cfg)

	if err := r.Run(cfg.Port()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
