package main

import (
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"web-installer-backend/boot"
	"web-installer-backend/global"
	router2 "web-installer-backend/router"
)

//go:embed dist/assets
var Resource embed.FS

//go:embed dist/index.html
var Html []byte
var port = "8200"

func main() {
	global.Asset = Resource
	global.Html = Html
	router := router2.NewRouter()
	gin.SetMode(gin.ReleaseMode)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		s.ListenAndServe()
	}()
	boot.Open(port)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
