package main

import (
	"github.com/poportss/finport-backend/internal/di"
	"github.com/poportss/finport-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega .env se existir
	_ = godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false, // já está ok para evitar prepared statements duplicados
	})
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao acessar sql.DB:", err)
	}

	// Ping para testar a conexão!
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Não foi possível conectar no banco de dados:", err)
	} else {
		log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	}

	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Inicializa servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	autoMigrate(db)

	container := di.NewContainer(db)
	di.SetupRoutes(router, container)

	routes := router.Routes()

	// Ordena rotas por path (alfabeticamente)
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})

	for _, route := range routes {
		log.Printf("🔗 %s %s", route.Method, route.Path)
	}

	log.Println("Servidor iniciado na porta", port)

	router.Run(":" + port)

}

func autoMigrate(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.User{},
		&models.Brokerage{},
		&models.BrokerageAsset{},
		&models.BrokerageNote{},
		&models.Trade{},
		&models.WalletType{},
		&models.Wallet{},
		&models.SchemaVersion{}, // inclua sua nova model aqui também!
	)
	if err != nil {
		return err
	}

	return db.AutoMigrate(
		&models.SchemaVersion{}, // inclua aqui!
	)

	return nil
}
