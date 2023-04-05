package repository

import (
	"context"
	"fmt"
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

func (postgresDbRepo *PostgresDbRepo) IsRoomAvailable(roomId int, startDate time.Time, endDate time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	var rowAmount int

	const query = `
				select count(id)
				from room_restrictions
				where
				    room_id = $1 
				    $2 >= start_date and $3 <= end_date
	`
	row := postgresDbRepo.SqlDB.QueryRowContext(ctx, query, roomId, startDate, endDate)
	err = row.Scan(&rowAmount)
	if err != nil {
		fmt.Println("IF ERROR :", err)
		return false, err
	}

	fmt.Println("row :", rowAmount)

	return !(rowAmount > 0), err
}

func (postgresDbRepo *PostgresDbRepo) GetAllAvailableRooms(startDate, endDate time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	const query = `
			select
				r.id, r.room_name
			from
			    rooms r
			where r.id not in (
								select
								    room_id from room_restrictions rr
								where
								    $1 >= start_date and $2 <= end_date
			            		)
	`
	rows, err := postgresDbRepo.SqlDB.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room

		err = rows.Scan(&room.ID, &room.Name)
		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}

	// Рекомендуется проводить еще одну проверку на содержание ошибок в самих rows:
	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}
