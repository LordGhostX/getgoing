package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/model"
	"todo/views"
)

func create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			err := model.CreateTodo()
			if err != nil {
				writer.Write([]byte("Some error occurred!"))
				log.Fatal(err)
			} else {
				writer.WriteHeader(http.StatusCreated)
				data := views.Response{
					Code: http.StatusCreated,
					Body: "successfully created todo",
				}
				err := json.NewEncoder(writer).Encode(data)
				if err != nil {
					writer.Write([]byte("Some error occurred!"))
					log.Fatal(err)
				}
			}
		}
	}
}
