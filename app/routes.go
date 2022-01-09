package app

import (
	"net/http"
	"runtime"
	"strconv"
	"trygonic/app/utils/Logger"
	"trygonic/app/utils/response"

	"github.com/gin-gonic/gin"
)

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	Logger.Get().Info(frame.File + " " + strconv.Itoa(frame.Line) + " " + frame.Function)
}

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		Logger.Get().Info("Ping")
		response.Send(c, gin.H{
			"all": "ok",
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.GET("/ping/:id", func(c *gin.Context) {
		Logger.Get().Info("GET Ping ID " + c.Param("id"))
		response.Send(c, gin.H{
			"ok": c.Param("id"),
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/ping/:id", func(c *gin.Context) {
		Logger.Get().Info("POST Ping ID " + c.Param("id"))
		response.Send(c, gin.H{
			"ok": c.Param("id"),
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/validate", func(c *gin.Context) {
		Logger.Get().Info("validate")
		response.Send(c, gin.H{
			"results": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/suggestions/suggest", func(c *gin.Context) {
		Logger.Get().Info("Suggestions suggest")
		response.Send(c, gin.H{
			"results": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/suggestions/fetch", func(c *gin.Context) {
		Logger.Get().Info("Suggestions fetch")
		response.Send(c, gin.H{
			"rules": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.POST("/api/validate/mlid", func(c *gin.Context) {
		Logger.Get().Info("validate mlid")
		response.Send(c, gin.H{
			"results": []int{},
		}, http.StatusOK, http.StatusText(http.StatusOK))
	})

	engine.NoRoute(func(c *gin.Context) {
		Logger.Get().Info("Route not found")
		response.Send(c, gin.H{}, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})
}
