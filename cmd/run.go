package cmd

import (
	"net/http"
	"time"
	"webdav/lib"
	"webdav/routers"
)

func Run()  {
	s:=&http.Server{
		Addr: ":10001",
		Handler: routers.InitRouter(),
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	}
	err:=s.ListenAndServe()
	if err != nil {
		lib.Log().Panic("%v",err)
	}
}
