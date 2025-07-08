package middleware

import (
	"fmt"
	"github.com/poportss/finport-backend/internal/baseservice"
	"github.com/poportss/finport-backend/internal/models"
	"github.com/poportss/finport-backend/internal/pkg/auth"
	"gorm.io/gorm"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/dto"
)

type srv struct {
	*baseservice.BaseService
}

func SetupJWTMiddleware(jwtKey []byte, timeout, maxRefresh time.Duration, baseService *baseservice.BaseService) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "finport",
		Key:         jwtKey,
		Timeout:     timeout,
		MaxRefresh:  maxRefresh,
		IdentityKey: "id",

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			idStr, ok := claims["id"].(string)
			if !ok || idStr == "" {
				return nil
			}
			idUint, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				return nil
			}
			return &models.User{
				Model: gorm.Model{
					ID: uint(idUint),
				},
				Name:         "",
				Document:     "",
				Email:        "",
				PasswordHash: "",
			}
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			user, ok := data.(*models.User)
			if !ok || user == nil {
				return jwt.MapClaims{}
			}
			return jwt.MapClaims{
				"id":   fmt.Sprintf("%d", user.ID),
				"name": user.Name,
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login dto.Login
			if err := c.ShouldBindJSON(&login); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			authService := auth.NewService(baseService)
			user, err := authService.Authenticate(login)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return user, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			// Pode adicionar lógica aqui se quiser
			return true
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},

		TokenLookup: "header: Authorization, query: token",
	})
}

// Extração segura do userID como uint
func ExtractUserIDFromContext(c *gin.Context) (uint, error) {
	claims := jwt.ExtractClaims(c)
	userIDStr, ok := claims["id"].(string)
	if !ok || userIDStr == "" {
		return 0, fmt.Errorf("user id não encontrado ou inválido")
	}
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("user id inválido")
	}
	return uint(userID), nil
}
