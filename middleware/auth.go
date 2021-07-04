package middleware

import (
	"net/http"
	"webdav/lib"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			v, ok := lib.UserMap[username]
			if ok {
				if v.Password == password {
					c.Set("user", v)
					c.Next()
				} else {
					lib.Log().Error("用户名密码错误，用户：%v，密码：%v", username, password)
					c.AbortWithStatus(http.StatusUnauthorized)
				}
			} else {
				lib.Log().Error("未找到该用户，用户：%v，密码：%v", username, password)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
