package app

import (
	"github.com/gin-gonic/gin"
	"strings"
	"trygonic/app/config"
	"trygonic/app/middlewares"
	"trygonic/app/utils/Logger"
)

func init() {
	// Add any code that needs to run before main
	Logger.Get().Info("Initializing App: ENV: " + strings.ToUpper(config.Values.Env))
}

func Init() *gin.Engine {
	gin.DisableConsoleColor()
	if config.Values.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middlewares.RequestLogger())
	engine.Use(middlewares.ResponseLogger())

	// Register routes
	RegisterRoutes(engine)

	return engine
}
