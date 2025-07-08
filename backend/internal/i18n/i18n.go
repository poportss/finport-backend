package i18n

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type MessageAlias string

const (
	CreateUserAlias        MessageAlias = "MSG001"
	InvalidRequestAlias    MessageAlias = "MSG002"
	UserCreateFailedAlias  MessageAlias = "MSG004"
	InvalidToken           MessageAlias = "MSG005"
	ErrorToCreateBrokerage MessageAlias = "MSG006"
	ErrorParseData         MessageAlias = "MSG007"
	ErrorCreateUser        MessageAlias = "MSG008"
	DefaultWalletName      MessageAlias = "MSG009"
)

// Pega idioma do cabeçalho (ou pt por padrão)
func getLanguageFromContext(c *gin.Context) string {
	lang := strings.ToLower(strings.TrimSpace(c.GetHeader("Accept-Language")))

	switch {
	case strings.HasPrefix(lang, "en"):
		return "en"
	case strings.HasPrefix(lang, "es"):
		return "es"
	default:
		return "pt"
	}
}

// Busca mensagem traduzida por código
func GetMessage(c *gin.Context, code MessageAlias, args ...interface{}) string {
	lang := getLanguageFromContext(c)

	var template string
	switch lang {
	case "en":
		template = EnMessages[code]
	case "es":
		template = EsMessages[code]
	default:
		template = PtMessages[code]
	}

	if template == "" {
		return fmt.Sprintf("[mensagem não encontrada: %s]", code)
	}
	return fmt.Sprintf(template, args...)
}
