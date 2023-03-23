package handlers

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"github.com/AnonymFromInternet/Motel/internal/render"
	"log"
	"net/http"
)

type Repository struct {
	AppConfig *app.Config
}

var Repo *Repository

func CreateNewRepository(appConfig *app.Config) *Repository {
	return &Repository{
		AppConfig: appConfig,
	}
}

func AreAskingToGetTheRepository(repository *Repository) {
	Repo = repository
}

func (repository *Repository) GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "main.page.gohtml"
	err := render.Template(writer, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'main'")
	}
}

func (repository *Repository) GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "contacts.page.gohtml"
	testData := make(map[string]interface{})
	testData["testData"] = "Test Data"

	err := render.Template(writer, fileName, &models.TemplatesData{BasicData: testData})

	if err != nil {
		log.Println("cannot render template with name 'contacts'")
	}
}
