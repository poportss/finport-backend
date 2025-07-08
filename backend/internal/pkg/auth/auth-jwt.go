package auth

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/poportss/finport-backend/internal/dto"
	"github.com/poportss/finport-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func (s srv) Authenticate(login dto.Login) (*models.User, error) {
	var user *models.User
	switch login.Provider {
	case "email":
		err := s.DB.Where("email = ?", login.Email).First(&user).Error
		if err != nil {
			return nil, errors.New("usuário não encontrado")
		}

		if user == nil {
			return nil, jwt.ErrFailedAuthentication
		}

		if !CheckPassword(user.PasswordHash, login.Password) {
			return nil, errors.New("senha inválida")
		}

		return user, nil
	default:
		return nil, errors.New("provider inválido")
	}
}

func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func HashPassword(password string) string {
	c, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Printf("encription failed: %v", err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s", c)
}
