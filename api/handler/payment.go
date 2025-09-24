package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/mp"
	"github.com/arthu/shop-api-go/internal/repo"
	"github.com/arthu/shop-api-go/internal/models"
)

// Mirrors Node route: POST /payment/notification
// Expects: query param orderID, body with data.id (payment id)
func RegisterPayment(r *gin.Engine, cfg *config.Config) {
	r.POST("/payment/notification", func(c *gin.Context) {
		orderID := c.Query("orderID")
		var body struct { Data struct{ ID string `json:"id"` } `json:"data"` }
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
			return
		}
		if orderID != "" && body.Data.ID != "" && cfg.MPAccessToken != "" {
			if status, err := mp.ClientImpl.GetPayment(cfg.MPAccessToken, body.Data.ID); err == nil && status != "" {
				// map MP status to our enum; keep same string if matches
				s := models.OrderStatus(status)
				_, _ = repo.UpdateOrderStatus(orderID, s)
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
