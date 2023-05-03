package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
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
	dbConnectionPool, err := prepareAppDataBeforeRunGetDBConnectionPool()
	if err != nil {
		log.Fatal("cannot prepare app data before server starting")
	}
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			log.Fatal("cannot close db connection pool")
		}
	}(dbConnectionPool.SQL)

	defer close(appConfig.MailChan)

	listenForMail()
	fmt.Println("Mail listener successfully started")

	server := &http.Server{
		Addr:    portNumber,
		Handler: getHandler(), // &appConfig в него передавать необязательно
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("err :", err)
		log.Fatal("cannot start server")
	}
}

func prepareAppDataBeforeRunGetDBConnectionPool() (*driver.DataBaseConnectionPool, error) {
	// Register in global encoding decoding for working with Session
	gob.Register(models.Admin{})
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register([]models.Room{})

	var err error

	mailChan := make(chan models.MailData)
	appConfig.MailChan = mailChan

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
		log.Fatal("[main]:[prepareAppDataBeforeRunGetDBConnectionPool] - cannot get dataBaseConnectionPool")
	}

	mainRepository := handlers.GetMainRepository(&appConfig, dataBaseConnectionPool)

	handlers.SetPkgRepoVariable(mainRepository)
	render.AsksToGet(&appConfig)
	helpers.AreAskingToGet(&appConfig)

	return dataBaseConnectionPool, nil
}
