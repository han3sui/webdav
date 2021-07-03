package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"webdav/lib"
)

type ErrorInfo struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Request   string `json:"request"`
	Timestamp int64  `json:"timestamp"`
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorInfo{
			Code:      404,
			Message:   "404 NOT FOUND",
			Request:   c.Request.Method + " " + c.Request.RequestURI,
			Timestamp: time.Now().Unix(),
		})
	}
}

func NotMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorInfo{
			Code:      404,
			Message:   "NOT ALLOWED METHOD",
			Request:   c.Request.Method + " " + c.Request.RequestURI,
			Timestamp: time.Now().Unix(),
		})
	}
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorInfo{
					Code:      500,
					Message:   "SERVER ERROR",
					Request:   c.Request.Method + " " + c.Request.RequestURI,
					Timestamp: time.Now().Unix(),
				})
				lib.Log().Panic("%v", err)
			}
		}()
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//lib.Log().Info("%v %v %v", c.Writer.Status(), c.Request.Method, c.Request.RequestURI)
	}
}
