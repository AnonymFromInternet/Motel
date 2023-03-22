package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"log"
	"net/http"
)

var appConfiguration *app.Config

func Template(writer http.ResponseWriter, templateFirstName string) error {
	templateCache, templateExistsInCache := appConfiguration.TemplatesCache[templateFirstName]

	if templateExistsInCache {
		// Хорошей практикой является использовать буфер, и только потом Execute для более точного отлова ошибок
		buffer := new(bytes.Buffer)

		err := templateCache.Execute(buffer, nil)
		if err != nil {
			fmt.Println(err)

			return err
		}

		_, err = buffer.WriteTo(writer)
		if err != nil {
			return err
		}
	} else {
		log.Fatal(errors.New("false file name! template for this file name cannot be exist"))
	}

	return nil
}

func AsksToGetTheAppConfig(config *app.Config) {
	appConfiguration = config
}
