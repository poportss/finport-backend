package user

import (
	"github.com/poportss/finport-backend/internal/dto"
	"github.com/poportss/finport-backend/internal/models"
)

func (s *srv) CreateUser(userRequest dto.CreateUserRequest) (*models.User, error) {

	user := &models.User{
		Name:         userRequest.Name,
		Document:     userRequest.Document,
		Email:        userRequest.Email,
		PasswordHash: userRequest.Password,
	}

	if err := s.DB.Model(&models.User{}).Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
