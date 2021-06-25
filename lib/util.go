package lib

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

func CreatePath(path string) (err error) {
	if !PathExists(path) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return
}

func PathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	return !os.IsNotExist(err)
}

func FormToBody(form map[string]string) (body io.Reader) {
	formData := url.Values{}
	for k, v := range form {
		formData.Add(k, v)
	}
	body = strings.NewReader(formData.Encode())
	return
}

func MultipartToBody(form map[string]string) (writer *multipart.Writer, body io.Reader) {
	buf := &bytes.Buffer{}
	writer = multipart.NewWriter(buf)
	for name, value := range form {
		_ = writer.WriteField(name, value)
	}
	_ = writer.Close()
	body = buf
	return
}
