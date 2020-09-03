package app

import (
    "github.com/gin-gonic/gin"
    "runtime"
    "strconv"
    "trygonic/app/utils/Logger"
    "trygonic/app/utils/response"
)

func trace2() {
    pc := make([]uintptr, 15)
    n := runtime.Callers(2, pc)
    frames := runtime.CallersFrames(pc[:n])
    frame, _ := frames.Next()
    Logger.Get().Info( frame.File + " " + strconv.Itoa(frame.Line) + " " + frame.Function)
}

func RegisterRoutes(engine *gin.Engine) {
    engine.GET("/ping/:id", func(c *gin.Context) {
        Logger.Get().Info("This is INFO")
        response.Send(c, gin.H{
            "a": "b",
        }, 200, "Success")
    })

    engine.POST("/ping/:id", func(c *gin.Context) {
        Logger.Get().Info("This is INFO")
        response.Send(c, gin.H{
            "a": "b",
        }, 200, "Success")
    })

    engine.POST("/api/validate", func(c *gin.Context) {
        Logger.Get().Info("This is INFO")
        response.Send(c, gin.H{
            "result": []int{},
        }, 200, "Success")
    })
}
