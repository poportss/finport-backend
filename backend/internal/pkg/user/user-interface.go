package user

import (
	"github.com/poportss/finport-backend/internal/dto"
	"github.com/poportss/finport-backend/internal/models"
)

type UserService interface {
	CreateUser(userRequest dto.CreateUserRequest) (*models.User, error)
}
