package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		writer.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/go/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World Go!"))
	})

	err := http.ListenAndServe("127.0.0.1:3000", mux)
	if err != nil {
		return 
	}
}
