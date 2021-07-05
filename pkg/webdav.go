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
		}
	} else {
		lib.Log().Error("用户不存在：%v", user)
	}
}
