package handlers

import (
	"encoding/json"
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
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'main'")
	}
}

func (repository *Repository) GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "contacts.page.gohtml"
	testData := make(map[string]interface{})
	testData["testData"] = "Test Data"

	err := render.Template(writer, request, fileName, &models.TemplatesData{BasicData: testData})

	if err != nil {
		log.Println("cannot render template with name 'contacts'")
	}
}

func (repository *Repository) GetHandlerAboutPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "about.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'about'")
	}
}

func (repository *Repository) GetHandlerRoomPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "room.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'room'")
	}
}

func (repository *Repository) GetHandlerBlueRoomPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "blueRoom.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'blueRoom'")
	}
}

func (repository *Repository) GetHandlerAvailabilityPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "availability.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'availability'")
	}
}

func (repository *Repository) PostHandlerAvailabilityPage(writer http.ResponseWriter, request *http.Request) {
	var basicData map[string]interface{}

	startDate := request.Form.Get("start-date")
	endDate := request.Form.Get("end-date")

	basicData["startDate"] = startDate
	basicData["endDate"] = endDate

	const fileName = "availability.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{
		BasicData: basicData,
	})
	if err != nil {
		log.Println("cannot render template with name 'availability: [method:post]'")
	}
}

type jsonResponse struct {
	IsAvailable bool   `json:"isAvailable"`
	Message     string `json:"message"`
}

func (repository *Repository) PostHandlerAvailabilityPageJSON(writer http.ResponseWriter, request *http.Request) {
	var err error
	response := jsonResponse{IsAvailable: true, Message: "Available"}

	responseInJsonFormat, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		log.Println("cannot convert response data to JSON")
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(responseInJsonFormat)
	if err != nil {
		log.Println("cannot convert response data to JSON")
		return
	}
}

func (repository *Repository) GetHandlerReservationPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "reservation.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'reservation'")
	}
}

func (repository *Repository) GetHandlerReservationConfirmPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "reservationConfirm.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		log.Println("cannot render template with name 'reservationConfirm'")
	}
}
