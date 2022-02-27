package headerReturn

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const prefix = "/header/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

// HandleHeaderReturn nil
func HandleHeaderReturn(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Println("start handling header return request")

	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError(
			fmt.Sprintf("path %s must start "+
				"with %s",
				request.URL.Path, prefix))
	}

	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for key, vals := range request.Header {
		flag := false
		for _, val := range vals {
			if flag {
				writer.Header().Add(key, val)
			} else {
				writer.Header().Set(key, val)
			}
		}
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	envs := os.Environ()
	for _, env := range envs {
		if strings.HasPrefix(env, "VERSION") {
			writer.Header().Set("Version",
				strings.Split(env, "=")[1])
		}
	}

	fmt.Println("end handling header return request")

	return nil
}
