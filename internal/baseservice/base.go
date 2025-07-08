package baseservice

import (
	"gorm.io/gorm"
)

// BaseService define uma estrutura base para serviços
type BaseService struct {
	DB *gorm.DB
}

// NewBaseService cria um novo serviço base com a conexão do banco
func NewBaseService(db *gorm.DB) *BaseService {
	return &BaseService{DB: db}
}
