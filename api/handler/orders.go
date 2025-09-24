package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/mp"
	"github.com/arthu/shop-api-go/internal/repo"
	"github.com/arthu/shop-api-go/api/middleware"
)

type OrderProductReq struct {
	ID       int64 `json:"id"`
	Quantity int   `json:"quantity"`
}

type CreateOrderReq struct {
	UserID    int64            `json:"userId"`
	OrderDate time.Time        `json:"orderDate"`
	Products  []OrderProductReq `json:"products"`
}

func RegisterOrders(r *gin.Engine, cfg *config.Config) {
	admin := r.Group("")
	admin.Use(middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("admin"))
	admin.POST("/orders", func(c *gin.Context) {
		var req CreateOrderReq
		if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"message": "invalid body"}); return }

		// build product IDs
		ids := make([]int64, 0, len(req.Products))
		qty := make(map[int64]int)
		for _, p := range req.Products { ids = append(ids, p.ID); qty[p.ID] = p.Quantity }

		// fetch products from DB
		products, err := repo.GetProductsByIDs(ids)
		if err != nil || len(products) == 0 { c.JSON(http.StatusBadRequest, gin.H{"message": "Produtos inválidos"}); return }

		// create MP preference items
		items := make([]mp.PreferenceItem, 0, len(products))
		for _, p := range products {
			items = append(items, mp.PreferenceItem{ ID: strconv.FormatInt(p.ID, 10), Title: p.Name, Quantity: qty[p.ID], CurrencyID: "BRL", UnitPrice: p.Price })
		}
		orderID := uuid.NewString()
		pref := mp.PreferenceRequest{ Items: items, NotificationURL: cfg.MPNotificationURL + "/payment/notification?orderID=" + orderID }
		paymentURL, err := mp.ClientImpl.CreatePreference(cfg.MPAccessToken, pref)
		if err != nil || paymentURL == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar o pedido"}); return
		}

		order := models.Order{ ID: orderID, UserID: req.UserID, OrderDate: req.OrderDate, Status: models.OrderPending, PaymentURL: paymentURL }
		if err := repo.InsertOrder(&order); err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao persistir o pedido"}); return }

		var itemsToInsert []struct{ ProductID int64; Quantity int }
		for _, p := range req.Products { itemsToInsert = append(itemsToInsert, struct{ ProductID int64; Quantity int }{ ProductID: p.ID, Quantity: p.Quantity }) }
		if err := repo.InsertOrderProducts(orderID, itemsToInsert); err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao persistir itens do pedido"}); return }

		c.JSON(http.StatusCreated, gin.H{"payment_url": paymentURL})
	})

	support := r.Group("")
	support.Use(middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("suport"))
	support.GET("/orders", func(c *gin.Context) {
		orders, err := repo.ListOrders()
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar todos os pedidos"}); return }
		c.JSON(http.StatusOK, orders)
	})

	support.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		order, err := repo.GetOrderByID(id)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar o pedido"}); return }
		if order == nil { c.JSON(http.StatusNotFound, gin.H{"message": "Pedido não encontrado"}); return }
		products, _ := repo.GetProductsForOrder(order.ID)
		order.Products = products
		c.JSON(http.StatusOK, order)
	})

	auth := r.Group("")
	auth.Use(middleware.AuthenticateToken(cfg))
	auth.GET("/orders/user/:userId", func(c *gin.Context) {
		userID, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
		orders, err := repo.GetOrdersByUserID(userID)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar os pedidos do usuário"}); return }
		c.JSON(http.StatusOK, orders)
	})

	support.GET("/orders/product/:productId", func(c *gin.Context) {
		productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
		orders, err := repo.GetOrdersByProductID(productID)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar os pedidos do produto"}); return }
		c.JSON(http.StatusOK, orders)
	})

	admin.PATCH("/orders/:id/status", func(c *gin.Context) {
		id := c.Param("id")
		var body struct{ Status models.OrderStatus `json:"status"` }
		if err := c.ShouldBindJSON(&body); err != nil { c.JSON(http.StatusBadRequest, gin.H{"message": "invalid body"}); return }
		ok, err := repo.UpdateOrderStatus(id, body.Status)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar o status do pedido"}); return }
		if !ok { c.JSON(http.StatusNotFound, gin.H{"message": "Pedido não encontrado"}); return }
		c.JSON(http.StatusOK, gin.H{"message": "Status do pedido atualizado com sucesso"})
	})
}
