package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"webdav/lib"

	"golang.org/x/net/webdav"

	"github.com/gin-gonic/gin"
)

func InitWebdav(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		lib.Log().Error("用户不存在：%v", user)
		return
	}
	value, ok := user.(lib.UserInfo)
	if !ok {
		lib.Log().Error("用户解析失败：%v", user)
		return
	}
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
}

func ListDir(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		lib.Log().Error("用户不存在：%v", user)
		return
	}
	value, ok := user.(lib.UserInfo)
	if !ok {
		lib.Log().Error("用户解析失败：%v", user)
		return
	}
	path := value.Dir + c.Request.URL.Path
	files, err := ioutil.ReadDir(path)
	if err != nil {
		lib.Log().Error("目录读取失败：", err)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	w := c.Writer
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range files {
		name := d.Name()
		if d.IsDir() {
			name += "/"
			fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", name, name)
		} else {
			fmt.Fprintf(w, "<a href=\"%s\" download>%s</a>\n", name, name)
		}
	}
	fmt.Fprintf(w, "</pre>\n")
}
