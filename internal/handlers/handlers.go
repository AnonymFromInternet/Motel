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
	"log"
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
	const templateName = "main.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "contacts.page.gohtml"

	err := render.Template(writer, request, templateName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerAboutPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "about.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerRoomPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "room.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerBlueRoomPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "blueRoom.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetHandlerAvailabilityPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "availability.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{})

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

	const templateName = "availability.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{
		BasicData: basicData,
	})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

type jsonResponse struct {
	IsAvailable bool          `json:"isAvailable"`
	Rooms       []models.Room `json:"rooms"`
}

func (repository *Repository) PostHandlerAvailabilityPageJSON(writer http.ResponseWriter, request *http.Request) {
	var err error
	var isAvailable bool
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
	}
	repository.AppConfig.Session.Put(request.Context(), "rooms", rooms)

	dates.StartDate = startDate
	dates.EndDate = endDate

	repository.AppConfig.Session.Put(request.Context(), "dates", dates)

	response := jsonResponse{IsAvailable: isAvailable, Rooms: rooms}

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
	const templateName = "reservation.page.gohtml"
	basicData := make(map[string]interface{})
	basicData["rooms"] = repository.AppConfig.Session.Get(request.Context(), "rooms")
	basicData["dates"] = repository.AppConfig.Session.Get(request.Context(), "dates")

	err := render.Template(writer, request, templateName, &models.TemplatesData{
		BasicData: basicData,
	})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

// TempDataReservationConfirmPage // TODO перенести этот костыль в session
var TempDataReservationConfirmPage = make(map[string]interface{})

func (repository *Repository) PostHandlerReservationPage(writer http.ResponseWriter, request *http.Request) {
	dates, ok := repository.AppConfig.Session.Get(request.Context(), "dates").(models.Reservation)
	if !ok {
		helpers.LogServerError(writer, errors.New("[package handlers]:[funcPostHandlerReservationPage] - cannot cast data type"))
		return
	}

	roomName := request.Form.Get("chosen-room")
	roomId, err := repository.DataBaseRepoInterface.GetRoomIdBy(roomName)
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

	mailDataForAdmin := models.MailData{
		ClientName:    reservation.FirstName,
		ClientSurname: reservation.LastName,
		RoomName:      roomName,
		From:          reservation.Email,
		To:            "newReservation@com.com",
		Subject:       "New reservation",
		Content: fmt.Sprintf(
			"<p>New reservation was created from %s %s. Room is %s. Arrival: %s. Departure: %s </p>",
			reservation.FirstName,
			reservation.LastName,
			roomName,
			dates.StartDate.Format("2006-01-02"),
			dates.EndDate.Format("2006-01-02"),
		),
	}
	repository.AppConfig.MailChan <- mailDataForAdmin
}

func (repository *Repository) GetHandlerReservationConfirmPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "reservationConfirm.page.gohtml"
	err := render.Template(writer, request, templateName, &models.TemplatesData{
		// TODO перенести этот костыль в session
		BasicData: TempDataReservationConfirmPage,
	})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) GetLoginPage(writer http.ResponseWriter, request *http.Request) {
	const templateName = "login.page.gohtml"

	basicData := make(map[string]interface{})
	loginError := repository.AppConfig.Session.Get(request.Context(), "loginError")
	basicData["loginError"] = loginError

	err := render.Template(writer, request, templateName, &models.TemplatesData{
		BasicData: basicData,
	})
	if err != nil {
		helpers.LogServerError(writer, err)
	}
}

func (repository *Repository) PostLoginPage(writer http.ResponseWriter, request *http.Request) {
	// Good practice is to renew token everytime when a user makes login or logout
	err := repository.AppConfig.Session.RenewToken(request.Context())
	if err != nil {
		log.Fatal("[package handlers]:[PostLoginPage] - cannot renew token")
	}

	email := request.Form.Get("email")
	password := request.Form.Get("password")

	adminId, _, err := repository.DataBaseRepoInterface.AuthenticateGetAdminId(email, password)
	if err != nil {
		repository.AppConfig.Session.Put(request.Context(), "loginError", true)
		http.Redirect(writer, request, "/login", http.StatusSeeOther)

		return
	}

	admin, err := repository.DataBaseRepoInterface.GetAdminBy(adminId)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}

	err = repository.DataBaseRepoInterface.UpdateAdmin(admin)
	if err != nil {
		helpers.LogServerError(writer, err)
		return
	}

	repository.AppConfig.Session.Put(request.Context(), "adminEmail", admin.Email)

	http.Redirect(writer, request, "/main", http.StatusSeeOther)
}

func (repository *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	_ = repository.AppConfig.Session.Destroy(r.Context())
	_ = repository.AppConfig.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (repository *Repository) GetAdminDashboard(w http.ResponseWriter, r *http.Request) {
	const templateName = "adminDashboard.page.gohtml"
	_ = render.Template(w, r, templateName, &models.TemplatesData{})
}

func (repository *Repository) GetAdminClientsReservations(w http.ResponseWriter, r *http.Request) {
	const templateName = "clients-reservations.page.gohtml"

	clientsReservations, err := repository.DataBaseRepoInterface.GetClientsOrAdminsReservations(helpers.RestrictionTypeReservation)
	if err != nil {
		helpers.LogServerError(w, err)
		return
	}

	basicData := make(map[string]interface{})
	basicData["clientsReservations"] = clientsReservations

	err = render.Template(w, r, templateName, &models.TemplatesData{
		BasicData: basicData,
	})
	if err != nil {
		helpers.LogServerError(w, err)
	}
}

func (repository *Repository) GetAdminAdminsReservations(w http.ResponseWriter, r *http.Request) {
	const templateName = "admins-reservations.page.gohtml"

	//adminsReservations, err := repository.DataBaseRepoInterface.GetClientsOrAdminsReservations(helpers.RestrictionTypeService)
	//if err != nil {
	//	helpers.LogServerError(w, err)
	//	return
	//}
	//

	// происходит какая то жопа в мидлвере, так как если не логиниться, то все будет работать нормально
	basicData := make(map[string]interface{})
	basicData["adminsReservations"] = "adminsReservations"

	_ = render.Template(w, r, templateName, &models.TemplatesData{
		BasicData: basicData,
	})
}

func (repository *Repository) GetAdminReservationsCalendar(w http.ResponseWriter, r *http.Request) {
	const templateName = "reservations-calendar.page.gohtml"
	_ = render.Template(w, r, templateName, &models.TemplatesData{})
}
