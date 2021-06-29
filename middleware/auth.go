package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webdav/lib"
)

//type LoginInfo struct {
//	User string `form:''`
//}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			lib.Log().Info("用户名：%v，密码：%v", username, password)
			c.Next()
		}
	}
}
