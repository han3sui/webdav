package pkg

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"net/http"
	"webdav/lib"
)

func InitWebdav(c *gin.Context) {
	user, ok := c.Get("user")
	if ok {
		value, ok1 := user.(lib.UserInfo)
		if ok1 {
			lib.Log().Info("当前用户%v", value)
			for _, v := range value.Dir {
				fs := &webdav.Handler{
					Prefix:     c.Request.URL.Path,
					FileSystem: webdav.Dir(v),
					LockSystem: webdav.NewMemLS(),
					Logger: func(r *http.Request, err error) {
						if err != nil {
							lib.Log().Error("%v", err)
						} else {
							lib.Log().Info("Webdav Log： %s: %s", r.Method, r.URL)
						}
					},
				}
				fs.ServeHTTP(c.Writer, c.Request)
			}
		}
	} else {
		lib.Log().Error("用户不存在：%v", user)
	}
}
