package main

import (
	"database/sql"
	"encoding/gob"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/driver"
	"github.com/AnonymFromInternet/Motel/internal/handlers"
	"github.com/AnonymFromInternet/Motel/internal/helpers"
	"github.com/AnonymFromInternet/Motel/internal/models"
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
	dataBaseConnectionPool, err := prepareAppDataBeforeRun()
	if err != nil {
		log.Fatal("cannot prepare app data before server starting")
	}
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			log.Fatal("cannot close db connection pool")
		}
	}(dataBaseConnectionPool.SQL)

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
func prepareAppDataBeforeRun() (*driver.DataBaseConnectionPool, error) {
	// Вроде это не обязательно
	gob.Register(models.User{})
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	var err error

	appConfig.TemplatesCache, err = templatesCache.Create()
	if err != nil {
		log.Fatal("[package main]:[func main] - cannot get app config")
		return nil, err
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

	// Connection to the database
	dataBaseConnectionPool, err := driver.GetDataBaseConnectionPool("host=localhost port=5432 dbname=Motel user=arturkeil password=")
	if err != nil {
		log.Fatal("[main]:[prepareAppDataBeforeRun] - cannot get dataBaseConnectionPool")
	}

	repositoryFromHandlersPkg := handlers.CreateNewRepository(&appConfig, dataBaseConnectionPool)

	handlers.AreAskingToGetTheRepository(repositoryFromHandlersPkg)
	render.AsksToGetTheAppConfig(&appConfig)
	helpers.AreAskingToGetAppConfig(&appConfig)

	return dataBaseConnectionPool, nil
}
