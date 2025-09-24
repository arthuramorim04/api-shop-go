package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/models"
)

// Handler exposes methods used by API layer to bind endpoints.
type Handler interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct{ svc Service }

func NewHandler(svc Service) Handler { return &handler{svc: svc} }

func (h *handler) List(c *gin.Context) {
	items, err := h.svc.List()
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar produtos"}); return }
	c.JSON(http.StatusOK, items)
}

func (h *handler) Create(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
	id, err := h.svc.Create(&p)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"}); return }
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *handler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
	p, err := h.svc.GetByID(id)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar produto"}); return }
	if p == nil { c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"}); return }
	c.JSON(http.StatusOK, p)
}

func (h *handler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
	ok, err := h.svc.Update(id, &p)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"}); return }
	if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"}); return }
	c.Status(http.StatusOK)
}

func (h *handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
	ok, err := h.svc.Delete(id)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir produto"}); return }
	if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"}); return }
	c.Status(http.StatusNoContent)
}
