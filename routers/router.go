package routers

import (
	"net/http"
	"webdav/middleware"
	"webdav/pkg"

	"github.com/gin-gonic/gin"
)

func InitRouter() http.Handler {
	r := gin.New()
	r.NoMethod(middleware.NotMethod())
	r.NoRoute(middleware.NotFound())
	r.Use(middleware.Cors(), middleware.Logger(), middleware.Recover())
	v1 := r.Group("/")
	v1.Use(middleware.Auth())
	{
		v1.Handle("PROPFIND", "/", pkg.InitWebdav)
	}
	return r
}
