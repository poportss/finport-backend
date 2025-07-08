package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/baseservice"
	"net/http"
)

type srv struct {
	*baseservice.BaseService
}

func NewService(base *baseservice.BaseService) *srv {
	return &srv{BaseService: base}
}

func ConfigureRoutes(r *gin.Engine, base *baseservice.BaseService, jwtMiddleware *jwt.GinJWTMiddleware) {
	//err := migrations.Migrate(base.DB, "auth", migration.Versions())
	//if err != nil {
	//	panic("❌ Erro ao rodar as migrations: " + err.Error())
	//}

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", jwtMiddleware.LoginHandler)
		authRoutes.GET("/refresh_token", jwtMiddleware.RefreshHandler)

		authRoutes.Use(jwtMiddleware.MiddlewareFunc())
		{
			authRoutes.GET("/validate", func(c *gin.Context) {
				claims := jwt.ExtractClaims(c)
				c.JSON(http.StatusOK, gin.H{"claims": claims})
			})
		}
	}

}
