package pkg
import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"net/http"
	"webdav/lib"
)
func InitWebdav(c *gin.Context)  {
	fs := &webdav.Handler{
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				lib.Log().Error("%v",err)
			}else {
				lib.Log().Info("Webdav %s: %s",r.Method,r.URL)
			}
		},
	}
	fs.ServeHTTP(c.Writer,c.Request)
}
