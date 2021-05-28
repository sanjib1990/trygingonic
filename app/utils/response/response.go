package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(c *gin.Context, T interface{}, code int, message string) {
	if code < 300 {
		code = http.StatusOK
	}

	if len(message) == 0 {
		message = "Success"
	}

	c.JSON(code, gin.H{
		"status"	: http.StatusText(code),
		"message"	: message,
		"data"		: T,
		"code"		: code,
	})
}
