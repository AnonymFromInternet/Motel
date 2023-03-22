package handlers

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/render"
	"log"
	"net/http"
)

type Repository struct {
	AppConfig *app.Config
}

var repo *Repository

func CreateNewRepository(appConfig *app.Config) *Repository {
	return &Repository{
		AppConfig: appConfig,
	}
}

func AsksToGetTheRepository(repository *Repository) {
	repo = repository
}

func (repository *Repository) GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "main.page.gohtml"
	err := render.Template(writer, fileName)

	if err != nil {
		log.Println("cannot render template with name 'main'")
	}
}

func (repository *Repository) GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "contacts.page.gohtml"
	err := render.Template(writer, fileName)

	if err != nil {
		log.Println("cannot render template with name 'contacts'")
	}
}
