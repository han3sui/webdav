package middleware

import (
	"net/http"
	"webdav/lib"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// OPTIONS 请求不需要鉴权，否则Windows10下无法保存文档
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		// 开始校验用户密码
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		v, ok := lib.UserMap[username]
		if !ok {
			lib.Log().Error("未找到该用户，用户：%v，密码：%v", username, password)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if v.Password != password {
			lib.Log().Error("用户名密码错误，用户：%v，密码：%v", username, password)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if v.Readonly && c.Request.Method != "PROPFIND" && c.Request.Method != "GET" {
			lib.Log().Error("【%v】当前用户无操作权限(%v)", username, c.Request)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Set("user", v)
		c.Next()
	}
}
