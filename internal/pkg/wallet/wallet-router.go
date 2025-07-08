package wallet

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/baseservice"
	"github.com/poportss/finport-backend/internal/migrations"
	"github.com/poportss/finport-backend/internal/pkg/wallet/migration"
)

type srv struct {
	*baseservice.BaseService
}

func NewService(base *baseservice.BaseService) *srv {
	err := migrations.Migrate(base.DB, "user", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}

	return &srv{BaseService: base}
}

func ConfigureRoutes(r *gin.Engine, base *baseservice.BaseService, jwtMiddleware *jwt.GinJWTMiddleware) {
	service := NewService(base)
	api := r.Group("/api")

	walletRoutes := api.Group("/wallet")

	walletRoutes.Use(jwtMiddleware.MiddlewareFunc())
	{
		walletRoutes.POST("/createWallet", service.CreateWalletHandler)
	}

}
