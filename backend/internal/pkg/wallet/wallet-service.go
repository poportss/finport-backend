package wallet

import (
	"github.com/poportss/finport-backend/internal/dto"
	"github.com/poportss/finport-backend/internal/models"
)

func (s *srv) CreateWallet(createWalletRequest dto.CreateWalletRequest, userID uint) error {

	wallet := &models.Wallet{
		UserID:       userID,
		Name:         createWalletRequest.Name,
		WalletTypeID: createWalletRequest.WalletTypeID,
	}

	if err := s.DB.Model(&models.Wallet{}).Create(&wallet).Error; err != nil {
		return err
	}

	return nil
}
