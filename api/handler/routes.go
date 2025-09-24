package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/api/middleware"
	usersModule "github.com/arthu/shop-api-go/internal/modules/users"
	productsModule "github.com/arthu/shop-api-go/internal/modules/products"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	// Health
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	// Auth
	RegisterAuth(r, cfg)

	// Modules wiring
	// Users module (protected: support role)
	{
		urepo := usersModule.NewRepository()
		usvc := usersModule.NewService(urepo)
		uh := usersModule.NewHandler(usvc)

		users := r.Group("")
		users.Use(middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("suport"))
		users.GET("/users", uh.List)
		users.POST("/users", uh.Create)
		users.GET("/users/:id", uh.Get)
		users.PUT("/users/:id", uh.Update)
		users.DELETE("/users/:id", uh.Delete)
	}

	// Products module
	{
		prepo := productsModule.NewRepository()
		psvc := productsModule.NewService(prepo)
		ph := productsModule.NewHandler(psvc)

		// public list
		r.GET("/products", ph.List)

		admin := r.Group("")
		admin.Use(middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("admin"))
		admin.POST("/products", ph.Create)
		admin.GET("/products/:productId", ph.Get)
		admin.PUT("/products/:productId", ph.Update)
		admin.DELETE("/products/:productId", ph.Delete)
	}

	// Orders
	RegisterOrders(r, cfg)

	// Payment webhook
	RegisterPayment(r, cfg)
}
