package main

import (
    "trygonic/app"
    "trygonic/app/config"
)

func main() {
    // Disable Console Color
    engine := app.Init()
    _ = engine.Run(":" + config.Values.Port)
}
