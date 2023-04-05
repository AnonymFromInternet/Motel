package helpers

import (
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"net/http"
	"runtime/debug"
	"time"
)

var appConfig *app.Config

func AreAskingToGet(appConfigParam *app.Config) {
	appConfig = appConfigParam
}

func LogClientError(writer http.ResponseWriter, err error, status int) {
	appConfig.ErrorLogger.Println("error :", err, "status", status)
	http.Error(writer, http.StatusText(status), status)
}

func LogServerError(writer http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrorLogger.Println("error :\n", trace)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func LogInfo(info string) {
	appConfig.InfoLogger.Println(info)
}

const (
	RestrictionTypeReservation int = 1
	RestrictionTypeService     int = 2
)

func GetDatesInTimeFormat(request *http.Request) (time.Time, time.Time, error) {
	startDateString := request.Form.Get("start-date")
	endDateString := request.Form.Get("end-date")
	// casting format
	layout := "2006-01-02" // В таком виде почему то приходит из формы, хотя на сайте данные идут в обратном порядке

	startDate, err := time.Parse(layout, startDateString)
	endDate, err := time.Parse(layout, endDateString)
	if err != nil {
		return startDate, endDate, err
	}

	return startDate, endDate, nil
}
