package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	err := http.ListenAndServe(":10001", &webdav.Handler{
		FileSystem: webdav.Dir("C:/"),
		LockSystem: webdav.NewMemLS(),
	})
	if err != nil {
		msg := fmt.Sprintf("启动失败：%v", err)
		panic(msg)
	}
}
