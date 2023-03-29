package repository

import "github.com/AnonymFromInternet/Motel/internal/models"

type DataBaseRepoInterface interface {
	InsertReservation(reservation models.Reservation) error
}
