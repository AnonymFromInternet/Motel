package repository

import (
	"context"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"time"
)

func (postgresDbRepo *PostgresDbRepo) InsertReservation(reservation models.Reservation) error {
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
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := postgresDbRepo.SqlDB.ExecContext(
		ctx,
		query,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.PhoneNumber,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
