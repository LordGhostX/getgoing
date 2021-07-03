package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/views"
)

func ping() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			body := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			writer.WriteHeader(http.StatusOK)
			err := json.NewEncoder(writer).Encode(body)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
