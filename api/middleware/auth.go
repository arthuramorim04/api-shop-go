package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/arthu/shop-api-go/internal/config"
)

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func AuthenticateToken(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
		if cl, ok := token.Claims.(*Claims); ok {
			c.Set("claims", cl)
		}
		c.Next()
	}
}

func AuthorizeRole(required string) gin.HandlerFunc {
    return func(c *gin.Context) {
        claimsAny, exists := c.Get("claims")
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
            return
        }
        claims := claimsAny.(*Claims)
        if claims.Role != required {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
            return
        }
        c.Next()
    }
}

func AuthorizeAny(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        claimsAny, exists := c.Get("claims")
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
            return
        }
        claims := claimsAny.(*Claims)
        for _, r := range roles {
            if claims.Role == r {
                c.Next()
                return
            }
        }
        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
    }
}
