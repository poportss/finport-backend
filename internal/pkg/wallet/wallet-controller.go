package wallet

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/dto"
	"github.com/poportss/finport-backend/internal/i18n"
	"github.com/poportss/finport-backend/internal/middleware"
	"github.com/poportss/finport-backend/internal/rest"
)

func (s *srv) CreateWalletHandler(c *gin.Context) {
	var walletRequest dto.CreateWalletRequest

	if err := c.ShouldBindJSON(&walletRequest); err != nil {
		rest.ResponseBadRequest(c, err, i18n.ErrorParseData)
		return
	}

	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		return
	}

	err = s.CreateWallet(walletRequest, userID)
	if err != nil {
		rest.ResponseInternalServerError(c, err, i18n.ErrorCreateUser)
		return
	}

	rest.ResponseDefaultSuccess(c, i18n.CreateUserAlias)
	return
}
