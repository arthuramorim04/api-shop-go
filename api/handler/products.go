package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/api/middleware"
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
)

func RegisterProducts(r *gin.Engine, cfg *config.Config) {
	// public list
	r.GET("/products", func(c *gin.Context) {
		items, err := repo.ListProducts()
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar produtos"}); return }
		c.JSON(http.StatusOK, items)
	})

	admin := r.Group("")
	admin.Use(middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("admin"))
	admin.POST("/products", func(c *gin.Context) {
		var p models.Product
		if err := c.ShouldBindJSON(&p); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
		id, err := repo.CreateProduct(&p)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"}); return }
		c.JSON(http.StatusCreated, gin.H{"id": id})
	})
	admin.GET("/products/:productId", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
		p, err := repo.GetProductByID(id)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar produto"}); return }
		if p == nil { c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"}); return }
		c.JSON(http.StatusOK, p)
	})
	admin.PUT("/products/:productId", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
		var p models.Product
		if err := c.ShouldBindJSON(&p); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
		ok, err := repo.UpdateProduct(id, &p)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"}); return }
		if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"}); return }
		c.Status(http.StatusOK)
	})
	admin.DELETE("/products/:productId", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
		ok, err := repo.DeleteProduct(id)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir produto"}); return }
		if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"}); return }
		c.Status(http.StatusNoContent)
	})
}
