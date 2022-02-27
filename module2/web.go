package main

import (
	"cloud-native-learn/module2/headerReturn"
	"cloud-native-learn/module2/healthz"
	"cloud-native-learn/module2/mockerror"
	"fmt"
	"net/http"
	"reflect"
)

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v\n", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
			// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
			// 是不是 status 为 0 的时候都返回 200 呢？
			fmt.Printf("request ip is：%s, status code is：%v\n",
				request.RemoteAddr,
				reflect.ValueOf(writer).Elem().FieldByName("status"))
		}()

		err := handler(writer, request)

		if err != nil {
			fmt.Printf("Error occurred "+
				"handling request: %s",
				err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusInternalServerError)
				return
			}
		}
	}
}

func main() {
	// 模拟正常请求，完成 1、2、3 题
	http.HandleFunc("/header/",
		errWrapper(headerReturn.HandleHeaderReturn))

	// 模拟异常请求，测试 3 题
	http.HandleFunc("/error/",
		errWrapper(mockerror.HandleHeaderReturn))

	// 完成第 4 题
	http.HandleFunc("/healthz/",
		errWrapper(healthz.HandleHeaderReturn))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
