package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"net/http"
	"webdav/lib"
)

func InitWebdav(c *gin.Context) {
	fmt.Println()
	fs := &webdav.Handler{
		Prefix:     c.Request.URL.Path,
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				lib.Log().Error("%v", err)
			} else {
				lib.Log().Info("Webdav %s: %s", r.Method, r.URL)
			}
		},
	}
	fs.ServeHTTP(c.Writer, c.Request)
}
