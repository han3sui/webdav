package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webdav/middleware"
	"webdav/pkg"
)

func InitRouter() http.Handler  {
	r:=gin.New()
	r.NoMethod(middleware.NotFound())
	r.NoRoute(middleware.NotFound())
	r.Use(middleware.Cors(),middleware.Logger(),middleware.Recover())
	r.GET("/dav",pkg.InitWebdav)
	return r
}
