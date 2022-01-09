package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var cfg map[string]string
var confMu = &sync.Mutex{}
var Values *config

func AppLogFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] - %s [%s] %s %s %d \"%s\" %s\n",
		param.TimeStamp.Format(time.RFC3339),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func get(key string, _default string) string {
	confMu.Lock()
	defer confMu.Unlock()
	if val, ok := cfg[key]; ok {
		return val
	}
	val, ok := os.LookupEnv(key)
	if !ok {
		return _default
	}
	return val
}

func init() {
	env, err := godotenv.Read()
	log.Println("Loading env...")
	if err != nil {
		log.Println(".env file not found")
	}

	if env == nil {
		env = make(map[string]string)
	}

	// Assigining env to config
	cfg = env

	// populate the cfg
	_env := get("ENV", "local")
	Values = &config{}

	Values.Env = _env
	Values.IsProd = _env == "prod" || _env == "release"
	Values.Port = get("PORT", "8080")
	Values.LogLevel = get("LOG_LEVEL", "INFO")
}

type config struct {
	Port     string
	Env      string
	IsProd   bool
	LogLevel string
}
