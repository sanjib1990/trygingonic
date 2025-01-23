package main

import (
	"fmt"
	"trygonic/app"
	"trygonic/app/config"
)

func main() {
	// Disable Console Color
	engine := app.Init()
	fmt.Println(config.Values.Port)
	_ = engine.Run(":" + config.Values.Port)
}
