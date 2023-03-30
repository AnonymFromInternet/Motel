package repository

import "github.com/AnonymFromInternet/Motel/internal/models"

type DataBaseRepoInterface interface {
	InsertReservationGetReservationId(reservation models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
}
