package handlers

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"testing"
)

var appConfig *app.Config

// Уже содержит внутри себя обработчик на путь /main
func TestHandlers(t *testing.T) {
	multiplexer := TESTGetMultiplexer(appConfig)
	if multiplexer == nil {
		t.Error("cannot get multiplexer")
	}
}

// Нужен
