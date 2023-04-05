package repository

import (
	"github.com/AnonymFromInternet/Motel/internal/models"
	"time"
)

type DataBaseRepoInterface interface {
	InsertReservationGetReservationId(reservation models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	IsRoomAvailable(startDate, endDate time.Time, roomId int) (bool, error)
}
