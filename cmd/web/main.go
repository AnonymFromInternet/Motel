package main

import (
	"github.com/AnonymFromInternet/Motel/internal/handlers"
	"log"
	"net/http"
)

const portNumber = "localhost:8080"

func main() {
	http.HandleFunc("/main", handlers.GetHandlerMainPage)
	http.HandleFunc("/contacts", handlers.GetHandlerContactsPage)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal("cannot start server. Error :", err)
	}
}
