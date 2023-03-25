package main

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/handlers"
	"github.com/AnonymFromInternet/Motel/internal/helpers"
	"github.com/AnonymFromInternet/Motel/internal/render"
	"github.com/AnonymFromInternet/Motel/internal/templatesCache"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = "localhost:8080"

var appConfig app.Config
var session *scs.SessionManager

func main() {
	err := prepareAppDataBeforeRun()
	if err != nil {
		log.Fatal("cannot prepare app data before server starting")
	}

	server := &http.Server{
		Addr:    portNumber,
		Handler: getHandler(&appConfig),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start server")
	}
}

// Скорее всего нужно будет переименовать
func prepareAppDataBeforeRun() error {
	var err error

	appConfig.TemplatesCache, err = templatesCache.Create()
	if err != nil {
		log.Fatal("[package main]:[func main] - cannot get app config")
		return err
	}
	appConfig.IsDevelopmentMode = true

	appConfig.InfoLogger = log.New(os.Stdout, "[INFO]:\n", log.Ldate|log.Ltime)
	appConfig.ErrorLogger = log.New(os.Stdout, "[ERROR]:\n", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 15 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = !appConfig.IsDevelopmentMode // false -> http / true -> https
	appConfig.Session = session

	repository := handlers.CreateNewRepository(&appConfig)

	handlers.AreAskingToGetTheRepository(repository)
	render.AsksToGetTheAppConfig(&appConfig)
	helpers.AreAskingToGetAppConfig(&appConfig)

	return nil
}
