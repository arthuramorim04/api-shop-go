package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
	"github.com/arthu/shop-api-go/internal/utils"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterAuth(r *gin.Engine, cfg *config.Config) {
	r.POST("/login", func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
			return
		}
		u, err := repo.GetUserByEmail(req.Email)
		if err != nil || u == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas"})
			return
		}
		if !utils.CheckPassword(u.Password, req.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas"})
			return
		}
		token, err := utils.GenerateJWT(cfg.JWTSecret, utils.JwtUser{ID: u.ID, Email: u.Email, Role: string(u.Role), Username: u.Username})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "login error"})
			return
		}
		resp := models.UserDTO{
			ID: u.ID,
			Username: u.Username,
			Email: u.Email,
			FirstName: u.FirstName,
			LastName: u.LastName,
			Address: u.Address,
			PhoneNumber: u.PhoneNumber,
			Role: u.Role,
			Token: token,
		}
		c.JSON(http.StatusOK, resp)
	})
}
