package user

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/baseservice"
	"github.com/poportss/finport-backend/internal/migrations"
	"github.com/poportss/finport-backend/internal/pkg/user/migration"
	"github.com/poportss/finport-backend/internal/pkg/wallet"
)

type srv struct {
	*baseservice.BaseService
	walletService wallet.WalletService // interface!

}

func NewService(base *baseservice.BaseService, walletService wallet.WalletService) *srv {
	err := migrations.Migrate(base.DB, "user", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}

	return &srv{BaseService: base, walletService: walletService}
}

func ConfigureRoutes(r *gin.Engine, base *baseservice.BaseService, jwtMiddleware *jwt.GinJWTMiddleware) {
	service := NewService(base, nil)
	api := r.Group("/api")

	userRoutes := api.Group("/user")

	userRoutes.POST("/createUser", service.CreateUserHandler)

	userRoutes.Use(jwtMiddleware.MiddlewareFunc())
	{
		//userRoutes.POST("/createUserTradeLink", service.CreateUserTradeLinkHandler)
	}

}
