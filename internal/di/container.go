package di

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/baseservice"
	"github.com/poportss/finport-backend/internal/middleware"
	"github.com/poportss/finport-backend/internal/pkg/auth"
	"github.com/poportss/finport-backend/internal/pkg/user"
	"github.com/poportss/finport-backend/internal/pkg/wallet"
	"gorm.io/gorm"
	"log"
	"time"
)

type Container struct {
	BaseService   *baseservice.BaseService
	JWTMiddleware *jwt.GinJWTMiddleware
	UserService   user.UserService
	WalletService wallet.WalletService
}

func NewContainer(db *gorm.DB) *Container {
	baseService := baseservice.NewBaseService(db)

	// Cria os serviços concretos, injetando o que precisar
	walletService := wallet.NewService(baseService)
	userService := user.NewService(baseService, walletService)

	// Configura o middleware JWT
	jwtMiddleware, err := middleware.SetupJWTMiddleware(
		[]byte("sua-jwt-secret"), time.Hour, time.Hour*24, baseService, // injete o serviço de auth real!
	)
	if err != nil {
		log.Fatalf("Erro ao configurar JWT: %v", err)
	}

	return &Container{
		BaseService:   baseService,
		JWTMiddleware: jwtMiddleware,
		UserService:   userService,
		WalletService: walletService,
	}
}

// Exemplo de setup das rotas fora do container (ou pode ser um método do container!)
func SetupRoutes(router *gin.Engine, c *Container) {
	auth.ConfigureRoutes(router, c.BaseService, c.JWTMiddleware)
	user.ConfigureRoutes(router, c.BaseService, c.JWTMiddleware)
	wallet.ConfigureRoutes(router, c.BaseService, c.JWTMiddleware)
}
