package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webdav/lib"
	"webdav/middleware"
	"webdav/pkg"
)

/**
LOCK。锁定资源，使用 Lock-Token: 标头。
UNLOCK。解除锁定，使用 Lock-Token: 标头。
PROPPATCH。设置、更改或删除单个资源的特性。
PROPFIND。用于获取一个或多个资源的一个或多个特性信息。该请求可能会包含一个值为 0、1或infinity的Depth: 标头。其中，0表示指定将获取指定URI处的集合的特性（也就是该文件或目录）；1表示指定将获取该集合以及位于该指定URI之下与其紧邻的资源的特性（非嵌套的子目录或子文件）；infinity表示指定将获取全部子目录或子文件（深度过大会加重对服务器的负担）。
COPY。复制资源，可以使用 Depth: 标头移动资源，使用 Destination: 标头指定目标。如果需要，COPY 方法也使用 Overwrite: 标头。
MOVE。移动资源，可以使用 Depth: 标头移动资源，使用 Destination: 标头指定目标。如果需要，MOVE 方法也使用 Overwrite: 标头。
MKCOL。用于创建新集合（对应目录）。
**/
//var method = []string{"LOCK", "UNLOCK", "PROPPATCH", "PROPFIND", "COPY", "MOVE", "MKCOL"}

func InitRouter() http.Handler {
	r := gin.New()
	r.NoMethod(middleware.NotMethod())
	r.NoRoute(middleware.NotFound())
	r.Use(middleware.Cors(), middleware.Logger(), middleware.Recover())
	v1 := r.Group(lib.Config.Server.Route)
	v1.Use(middleware.Auth())
	{
		v1.Any("/*path", pkg.InitWebdav)
		//v1.Any("", pkg.InitWebdav)
		//v1.GET("/*path", pkg.ListDir)
		//v1.Handle("PROPFIND", "/*path", pkg.InitWebdav)
		//v1.Handle("PROPFIND", "", pkg.InitWebdav)
		//v1.Handle("MKCOL", "/*path", pkg.InitWebdav)
		//v1.Handle("LOCK", "/*path", pkg.InitWebdav)
		//v1.Handle("UNLOCK", "/*path", pkg.InitWebdav)
		//v1.Handle("PROPPATCH", "/*path", pkg.InitWebdav)
		//v1.Handle("COPY", "/*path", pkg.InitWebdav)
		//v1.Handle("MOVE", "/*path", pkg.InitWebdav)
	}
	return r
}
