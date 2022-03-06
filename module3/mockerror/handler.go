package mockerror

import (
	"fmt"
	"net/http"
)

const prefix = "/error/"

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
	fmt.Println("start handling mock error request")

	panic("模拟错误")

	return nil
}
