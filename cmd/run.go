package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"webdav/lib"
	"webdav/routers"
)

func Run()  {
	if !lib.Config.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	s:=&http.Server{
		Addr: lib.Config.Server.Addr,
		Handler: routers.InitRouter(),
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	}
	err:=s.ListenAndServe()
	if err != nil {
		lib.Log().Panic("%v",err)
	}
}
