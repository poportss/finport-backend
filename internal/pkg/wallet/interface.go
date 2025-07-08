package wallet

import "github.com/poportss/finport-backend/internal/dto"

type WalletService interface {
	CreateWallet(createWalletRequest dto.CreateWalletRequest, userID uint) error
}
