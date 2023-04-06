package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/driver"
	"github.com/AnonymFromInternet/Motel/internal/helpers"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"github.com/AnonymFromInternet/Motel/internal/render"
	"github.com/AnonymFromInternet/Motel/internal/repository"
	repository2 "github.com/AnonymFromInternet/Motel/internal/repository/dbRepo"
	"net/http"
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

// PostHandlerAvailabilityPage TODO Данный хэндлер пока не используется и не факт что он вообще нужен будет
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
	IsAvailable bool          `json:"isAvailable"`
	Message     string        `json:"message"`
	Rooms       []models.Room `json:"rooms"`
}

func (repository *Repository) PostHandlerAvailabilityPageJSON(writer http.ResponseWriter, request *http.Request) {
	var err error
	var isAvailable bool
	var message string
	var dates models.Reservation

	startDate, endDate, err := helpers.GetDatesInTimeFormat(request)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}

	rooms, err := repository.DataBaseRepoInterface.GetAllAvailableRooms(startDate, endDate)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}
	if len(rooms) > 0 {
		isAvailable = true
		message = "Available"
	}
	repository.AppConfig.Session.Put(request.Context(), "rooms", rooms)

	dates.StartDate = startDate
	dates.EndDate = endDate

	repository.AppConfig.Session.Put(request.Context(), "dates", dates)

	response := jsonResponse{IsAvailable: isAvailable, Message: message, Rooms: rooms}

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
	basicData := make(map[string]interface{})
	basicData["rooms"] = repository.AppConfig.Session.Get(request.Context(), "rooms")
	basicData["dates"] = repository.AppConfig.Session.Get(request.Context(), "dates")

	err := render.Template(writer, request, fileName, &models.TemplatesData{
		BasicData: basicData,
	})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

// TempDataReservationConfirmPage // TODO перенести этот костыль в session
var TempDataReservationConfirmPage = make(map[string]interface{})

func (repository *Repository) PostHandlerReservationPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("PostHandlerReservationPage()")

	dates, ok := repository.AppConfig.Session.Get(request.Context(), "dates").(models.Reservation)
	if !ok {
		helpers.LogServerError(writer, errors.New("[package handlers]:[funcPostHandlerReservationPage] - cannot cast data type"))
		return
	}

	roomName := request.Form.Get("chosen-room")
	roomId, err := repository.DataBaseRepoInterface.GetRoomIdBy(roomName)
	fmt.Println("room id is :", roomId)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}

	reservation := models.Reservation{
		FirstName:   request.Form.Get("first-name"),
		LastName:    request.Form.Get("last-name"),
		Email:       request.Form.Get("email"),
		PhoneNumber: request.Form.Get("phone-number"),
		StartDate:   dates.StartDate,
		EndDate:     dates.EndDate,
		RoomId:      roomId,
	}

	reservationId, err := repository.DataBaseRepoInterface.InsertReservationGetReservationId(reservation)

	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}

	roomRestriction := models.RoomRestriction{
		StartDate:     dates.StartDate,
		EndDate:       dates.EndDate,
		RoomId:        roomId,
		ReservationId: reservationId,
		RestrictionId: helpers.RestrictionTypeReservation,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = repository.DataBaseRepoInterface.InsertRoomRestriction(roomRestriction)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}

	// TODO перенести этот костыль в session
	TempDataReservationConfirmPage["reservation"] = reservation
	TempDataReservationConfirmPage["sd"] = reservation.StartDate.Format("2006-01-02")
	TempDataReservationConfirmPage["ed"] = reservation.EndDate.Format("2006-01-02")

	http.Redirect(writer, request, "/reservation-confirm", http.StatusSeeOther)
}

func (repository *Repository) GetHandlerReservationConfirmPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "reservationConfirm.page.gohtml"
	err := render.Template(writer, request, fileName, &models.TemplatesData{
		BasicData: TempDataReservationConfirmPage,
	})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}
