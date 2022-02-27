package healthz

import (
	"fmt"
	"net/http"
)

const prefix = "/healthz/"

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
	fmt.Println("start handling healthz request")

	writer.Write([]byte("{\"message\": \"ok\"}"))

	fmt.Println("end handling healthz request")

	return nil
}
