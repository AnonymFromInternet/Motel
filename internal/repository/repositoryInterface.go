package repository

import (
	"github.com/AnonymFromInternet/Motel/internal/models"
	"time"
)

type DataBaseRepoInterface interface {
	InsertReservationGetReservationId(reservation models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	IsRoomAvailable(roomId int, startDate time.Time, endDate time.Time) (bool, error)
	GetAllAvailableRooms(startDate, endDate time.Time) ([]models.Room, error)
	GetRoomIdBy(roomName string) (int, error)
}
