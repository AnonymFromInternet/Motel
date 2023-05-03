package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"golang.org/x/crypto/bcrypt"
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
		return false, err
	}

	return !(rowAmount > 0), err
}

func (postgresDbRepo *PostgresDbRepo) GetAllAvailableRooms(startDate, endDate time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room
	const query = ` select r.id, r.room_name
     				from rooms r
     				    where r.id not in
					(select
						room_id
					 from room_restrictions rr
						where 
							($1 <= rr.start_date and $2 <= rr.end_date)
						or
							($1 <= rr.end_date and $2 >= rr.end_date)
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

func (postgresDbRepo *PostgresDbRepo) GetRoomIdBy(roomName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	var roomId int

	const query = `
				select id
				from rooms
				where
				    room_name = $1 
	`
	row := postgresDbRepo.SqlDB.QueryRowContext(ctx, query, roomName)
	err = row.Scan(&roomId)
	if err != nil {
		return roomId, err
	}

	return roomId, err
}

func (postgresDbRepo *PostgresDbRepo) GetAdminBy(id int) (models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	var admin models.Admin

	const query = `
				select id, first_name, last_name, email, access_level, password, created_at, updated_at
				from admins
				where
				    id = $1 
	`

	row := postgresDbRepo.SqlDB.QueryRowContext(ctx, query, id)
	err = row.Scan(
		&admin.ID,
		&admin.FirstName,
		&admin.LastName,
		&admin.Email,
		&admin.AccessLevel,
		&admin.Password,
		&admin.CreatedAt,
		&admin.UpdatedAt,
	)
	if err != nil {
		return admin, err
	}

	return admin, err
}

func (postgresDbRepo *PostgresDbRepo) UpdateAdmin(admin models.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	const query = `
				update admins set first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5 
	`

	_, err = postgresDbRepo.SqlDB.ExecContext(
		ctx,
		query,
		admin.FirstName,
		admin.LastName,
		admin.Email,
		admin.AccessLevel,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (postgresDbRepo *PostgresDbRepo) AuthenticateGetAdminId(email, testPassword string) (adminId int, hashedPassword string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const query = `select id, password from admins where email = $1`

	row := postgresDbRepo.SqlDB.QueryRowContext(ctx, query, email)
	err = row.Scan(&adminId, &hashedPassword)
	if err != nil {
		return adminId, hashedPassword, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return adminId, hashedPassword, errors.New("invalid password")
	} else if err != nil {
		return adminId, hashedPassword, err
	}

	return adminId, hashedPassword, err
}

func (postgresDbRepo *PostgresDbRepo) GetClientsOrAdminsReservations(restrictionType int) ([]models.Reservation, error) {
	fmt.Println("GetClientsOrAdminsReservations()")
	fmt.Println("------")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error
	var reservations []models.Reservation

	const query = `
						select res.id, res.first_name, res.last_name, res.email, res.phone, res.start_date, res.end_date, res.room_id,
						       res.created_at, res.updated_at
						from reservations res
						inner join room_restrictions rr on (res.id = rr.reservation_id)
						where rr.restriction_id = $1
	`

	rows, err := postgresDbRepo.SqlDB.QueryContext(ctx, query, restrictionType)

	if err != nil {
		return reservations, err
	}

	for rows.Next() {
		var reservation models.Reservation

		err = rows.Scan(
			&reservation.ID,
			&reservation.FirstName,
			&reservation.LastName,
			&reservation.Email,
			&reservation.PhoneNumber,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.RoomId,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, err
}
