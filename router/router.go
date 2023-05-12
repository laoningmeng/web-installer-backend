package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-installer-backend/app"
	"web-installer-backend/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/assets", http.FS(app.NewResource()))
	router.Use(middleware.Cors())
	group := router.Group("/api")
	{
		group.GET("/step/list", app.GetStepList)
		group.GET("/check/list", app.GetCheckList)
		group.POST("/config/save", app.SaveConfig)
		group.GET("/command", app.Exec)
		group.GET("/web", app.Resource)
	}
	return router
}
