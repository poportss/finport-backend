package dto

type CreateWalletRequest struct {
	Name         string `json:"name"`
	WalletTypeID uint   `json:"wallet_type_id"`
}
