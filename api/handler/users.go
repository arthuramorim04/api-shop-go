package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
	"github.com/arthu/shop-api-go/internal/utils"
)

func RegisterUsers(r *gin.RouterGroup) {
	r.GET("/users", func(c *gin.Context) {
		users, err := repo.ListUsers()
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"}); return }
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var u models.User
		if err := c.ShouldBindJSON(&u); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
		// hash password
		if u.Password != "" {
			h, err := utils.HashPassword(u.Password)
			if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "hash error"}); return }
			u.Password = h
		}
		id, err := repo.CreateUser(&u)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"}); return }
		u.ID = id
		c.JSON(http.StatusCreated, u)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		u, err := repo.GetUserByID(id)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o usuário"}); return }
		if u == nil { c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}); return }
		c.JSON(http.StatusOK, u)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var u models.User
		if err := c.ShouldBindJSON(&u); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "body inválido"}); return }
		ok, err := repo.UpdateUser(id, &u)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o usuário"}); return }
		if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}); return }
		u.ID = id
		c.JSON(http.StatusOK, u)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		ok, err := repo.DeleteUser(id)
		if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o usuário"}); return }
		if !ok { c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"}); return }
		c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
	})
}
