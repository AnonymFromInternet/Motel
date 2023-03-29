package helpers

import (
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"net/http"
	"runtime/debug"
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
