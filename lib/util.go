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

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)

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
