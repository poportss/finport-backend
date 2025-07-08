package user

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/dto"
	"github.com/poportss/finport-backend/internal/i18n"
	"github.com/poportss/finport-backend/internal/rest"
)

func (s *srv) CreateUserHandler(c *gin.Context) {
	var userRequest dto.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest.ResponseBadRequest(c, err, i18n.ErrorParseData)
		return
	}

	user, err := s.CreateUser(userRequest)
	if err != nil {
		rest.ResponseInternalServerError(c, err, i18n.ErrorCreateUser)
		return
	}

	walletData := dto.CreateWalletRequest{
		Name:         i18n.GetMessage(c, i18n.DefaultWalletName),
		WalletTypeID: 1,
	}

	err = s.walletService.CreateWallet(walletData, user.ID)
	if err != nil {
		rest.ResponseInternalServerError(c, err, i18n.ErrorCreateUser)
		return
	}

	rest.ResponseDefaultSuccess(c, i18n.CreateUserAlias)
	return
}
