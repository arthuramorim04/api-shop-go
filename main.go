package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/db"
	"github.com/arthu/shop-api-go/router"
)

func main() {
	// Load env
	_ = os.Setenv("GIN_MODE", "release")
	cfg := config.Load()

	// Connect DB
	if err := db.Connect(cfg.DBDSN); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	router.RegisterRoutes(r, cfg)

	if err := r.Run(cfg.Port()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
