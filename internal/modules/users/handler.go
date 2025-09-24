package users

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
	users, err := h.svc.List()
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"}); return }
	c.JSON(http.StatusOK, users)
}

func (h *handler) Create(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
	id, err := h.svc.Create(&u)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"}); return }
	u.ID = id
	c.JSON(http.StatusCreated, u)
}

func (h *handler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u, err := h.svc.GetByID(id)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o usuário"}); return }
	if u == nil { c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}); return }
	c.JSON(http.StatusOK, u)
}

func (h *handler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
	ok, err := h.svc.Update(id, &u)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o usuário"}); return }
	if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}); return }
	u.ID = id
	c.JSON(http.StatusOK, u)
}

func (h *handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ok, err := h.svc.Delete(id)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o usuário"}); return }
	if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}); return }
	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
