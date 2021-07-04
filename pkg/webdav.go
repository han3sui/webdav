package pkg

import (
	"webdav/lib"

	"github.com/gin-gonic/gin"
)

func InitWebdav(c *gin.Context) {
	user, ok := c.Get("user")
	if ok {
		value, ok1 := user.(lib.UserInfo)
		if ok1 {
			value.Fs.ServeHTTP(c.Writer, c.Request)
			//lib.Log().Info("当前用户%v", value)
			//fs := &webdav.Handler{
			//	Prefix:     lib.Config.Server.Route,
			//	FileSystem: webdav.Dir(value.Dir),
			//	LockSystem: webdav.NewMemLS(),
			//	Logger: func(r *http.Request, err error) {
			//		if err != nil {
			//			lib.Log().Error("%v", err)
			//		} else {
			//			lib.Log().Info("Webdav Log： %s: %s", r.Method, r.URL)
			//		}
			//	},
			//}
			//fs.ServeHTTP(c.Writer, c.Request)
		}
	} else {
		lib.Log().Error("用户不存在：%v", user)
	}
}
