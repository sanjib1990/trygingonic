package middlewares

import (
	"strconv"
	"strings"
	"trygonic/app/utils/Logger"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var str strings.Builder
		str.Write([]byte("[REQUEST RECEIVED]"))

		// append Method
		str.Write([]byte(" " + c.Request.Method))

		// append URI
		str.Write([]byte(": " + c.Request.URL.Path))

		// append any query string
		queries, _ := jsoniter.MarshalToString(c.Request.URL.Query())
		str.Write([]byte(" | [QUERY]: " + queries))

		// append path params:
		param, _ := jsoniter.MarshalToString(c.Params)
		str.Write([]byte(" | [PATH PARAMS]: " + param))

		// append body
		body, _ := c.GetRawData()
		var s strings.Builder
		s.Write(body)
		str.Write([]byte(" | [BODY]: " + s.String()))
		str.Write([]byte(" | [AGENT]: " + c.Request.UserAgent()))

		Logger.Get().Info(str.String())
		// before request
		c.Next()
	}
}

func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// access the status we are sending
		var str strings.Builder
		str.Write([]byte("[RESPONSE]"))
		str.Write([]byte(" [STATUS]: " + strconv.Itoa(c.Writer.Status())))
		Logger.Get().Info(str.String())
	}
}
