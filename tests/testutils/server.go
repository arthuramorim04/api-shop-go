package testutils

import (
	"database/sql"
	
	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/api/handler"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/database"
)

// NewServer initializes a Gin engine with routes registered and a provided *sql.DB (sqlmock in tests).
func NewServer(cfg *config.Config, db *sql.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	if db != nil {
		database.SetDB(db)
	}
	r := gin.New()
	handler.RegisterRoutes(r, cfg)
	return r
}
