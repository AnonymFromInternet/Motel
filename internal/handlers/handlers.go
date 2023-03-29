package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/driver"
	"github.com/AnonymFromInternet/Motel/internal/helpers"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"github.com/AnonymFromInternet/Motel/internal/render"
	"github.com/AnonymFromInternet/Motel/internal/repository"
	repository2 "github.com/AnonymFromInternet/Motel/internal/repository/dbRepo"
	"net/http"
	"strconv"
	"time"
)

type Repository struct {
	AppConfig             *app.Config
	DataBaseRepoInterface repository.DataBaseRepoInterface
}

var Repo *Repository

func GetMainRepository(appConfig *app.Config, dbConnPool *driver.DataBaseConnectionPool) *Repository {
	return &Repository{
		AppConfig:             appConfig,
		DataBaseRepoInterface: repository2.GetPostgresDbRepo(appConfig, dbConnPool.SQL),
	}
}

func SetPkgRepoVariable(repository *Repository) {
	Repo = repository
}

func (repository *Repository) GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "main.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "contacts.page.gohtml"
	testData := make(map[string]interface{})
	testData["testData"] = "Test Data"

	err := render.Template(writer, request, fileName, &models.TemplatesData{BasicData: testData})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerAboutPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "about.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerRoomPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "room.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerBlueRoomPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "blueRoom.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerAvailabilityPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "availability.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})

	if err != nil {
		helpers.LogServerError(writer, err)
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
		helpers.LogServerError(writer, err)
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
		helpers.LogServerError(writer, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(responseInJsonFormat)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}
}

func (repository *Repository) GetHandlerReservationPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "reservation.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) PostHandlerReservationPage(writer http.ResponseWriter, request *http.Request) {
	startDateString := request.Form.Get("start-date")
	endDateString := request.Form.Get("end-date")

	fmt.Println()

	// format casting
	layout := "2006-01-02"

	startDate, err := time.Parse(layout, startDateString)
	endDate, err := time.Parse(layout, endDateString)
	if err != nil {
		helpers.LogServerError(writer, err)
	}

	roomId, err := strconv.Atoi(request.Form.Get("room_id"))
	if err != nil {
		helpers.LogServerError(writer, err)
	}

	reservation := models.Reservation{
		ID:          0,
		FirstName:   request.Form.Get("first-name"),
		LastName:    request.Form.Get("last-name"),
		Email:       request.Form.Get("email"),
		PhoneNumber: request.Form.Get("phone-number"),
		StartDate:   startDate,
		EndDate:     endDate,
		RoomId:      roomId,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Room:        models.Room{},
	}

	err = repository.DataBaseRepoInterface.InsertReservation(reservation)
	if err != nil {
		helpers.LogServerError(writer, err)
	}

	fmt.Println("PostHandlerReservationPage")
}

func (repository *Repository) GetHandlerReservationConfirmPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "reservationConfirm.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}
