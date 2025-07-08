package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poportss/finport-backend/internal/i18n"
	"io/ioutil
	"net/http"
	"time"
)

// Sucesso padrão, pode receber um alias para mensagem se quiser personalizar
func ResponseDefaultSuccess(c *gin.Context, msgAlias i18n.MessageAlias, args ...interface{}) {
	msg := i18n.GetMessage(c, msgAlias, args...)
	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS", "message": msg})
}

// Internal Server Error com i18n
func ResponseInternalServerError(c *gin.Context, err error, msgAlias i18n.MessageAlias, args ...interface{}) {
	msg := i18n.GetMessage(c, msgAlias, args...)
	c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILURE", "errorMessage": msg, "error": err.Error()})
}

// Bad Request com i18n
func ResponseBadRequest(c *gin.Context, err error, msgAlias i18n.MessageAlias, args ...interface{}) {
	msg := i18n.GetMessage(c, msgAlias, args...)
	c.JSON(http.StatusBadRequest, gin.H{"status": "FAILURE", "errorMessage": msg, "error": err.Error()})
}

// Unauthorized com i18n
func ResponseUnauthorize(c *gin.Context, msgAlias i18n.MessageAlias, args ...interface{}) {
	msg := i18n.GetMessage(c, msgAlias, args...)
	c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILURE", "errorMessage": msg})
	c.Abort()
}

// Forbidden com i18n
func ResponseForbidden(c *gin.Context, msgAlias i18n.MessageAlias, args ...interface{}) {
	msg := i18n.GetMessage(c, msgAlias, args...)
	c.JSON(http.StatusForbidden, gin.H{"status": "FAILURE", "errorMessage": msg})
	c.Abort()
}

// Not Found com i18n
func ResponseNotFound(c *gin.Context, msgAlias i18n.MessageAlias, args ...interface{}) {
	msg := i18n.GetMessage(c, msgAlias, args...)
	c.JSON(http.StatusNotFound, gin.H{"status": "FAILURE", "errorMessage": msg})
	c.Abort()
}

// GetMethod (mantém igual, pois é utilitário)
func GetMethod(url string) ([]byte, error) {
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("error in response: %s", resBody)
	}

	return resBody, nil
}
