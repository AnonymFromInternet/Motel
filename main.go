package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "Main page")

		if err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("cannot start server. Error :", err)
	}
}
