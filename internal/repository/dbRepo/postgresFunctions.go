package repository

import (
	"context"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"time"
)

func (postgresDbRepo *PostgresDbRepo) InsertReservationGetReservationId(reservation models.Reservation) (int, error) {
	const query = `insert into reservations (
			                          first_name,
			                          last_name,
			                          email,
			                          phone,
			                          start_date,
			                          end_date,
			                          room_id,
			                          created_at,
			                          updated_at
			                          )
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id
`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservationId int

	row := postgresDbRepo.SqlDB.QueryRowContext(
		ctx,
		query,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.PhoneNumber,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomId,
		reservation.CreatedAt,
		reservation.UpdatedAt,
	)

	err := row.Scan(&reservationId)
	if err != nil {
		return reservationId, err
	}

	// Проверить сам скан на наличие ошибок

	return reservationId, nil
}

func (postgresDbRepo *PostgresDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const query = `
			insert into room_restrictions (
			                               start_date,
			                               end_date,
			                               room_id,
			                               reservation_id,
			                               restriction_id,
			                               created_at,
			                               updated_at
			)
			values ($1, $2, $3, $4, $5, $6, $7)
`
	_, err := postgresDbRepo.SqlDB.ExecContext(ctx, query,
		roomRestriction.StartDate,
		roomRestriction.EndDate,
		roomRestriction.RoomId,
		roomRestriction.ReservationId,
		roomRestriction.RestrictionId,
		roomRestriction.CreatedAt,
		roomRestriction.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
