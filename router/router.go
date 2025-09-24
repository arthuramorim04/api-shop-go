package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/middleware"
	"github.com/arthu/shop-api-go/web"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	// Health
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	// Auth
	web.RegisterAuth(r, cfg)

	// Users
	users := r.Group("")
	users.Use(middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("suport"))
	web.RegisterUsers(users)

	// Products
	web.RegisterProducts(r, cfg)

	// Orders
	web.RegisterOrders(r, cfg)

	// Payment webhook
	web.RegisterPayment(r, cfg)
}
