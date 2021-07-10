package pkg

import (
	"net/http"
	"webdav/lib"

	"golang.org/x/net/webdav"

	"github.com/gin-gonic/gin"
)

func InitWebdav(c *gin.Context) {
	user, ok := c.Get("user")
	if ok {
		value, ok1 := user.(lib.UserInfo)
		if ok1 {
			fs := &webdav.Handler{
				Prefix:     lib.Config.Server.Route,
				FileSystem: webdav.Dir(value.Dir),
				LockSystem: webdav.NewMemLS(),
				Logger: func(request *http.Request, err error) {
					if err != nil {
						lib.Log().Error("【%v】%v", value.Name, err)
						return
					}
					lib.Log().Info("【%v】%v %v", value.Name, request.Method, request.URL)
				},
			}
			fs.ServeHTTP(c.Writer, c.Request)
			//value.Fs.ServeHTTP(c.Writer, c.Request)
		}
	} else {
		lib.Log().Error("用户不存在：%v", user)
	}
}
