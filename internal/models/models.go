package models

import "time"

type TemplatesData struct {
	BasicData   map[string]interface{}
	Error       string
	Warning     string
	CSRFToken   string
	ShowMessage string
}

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Room struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restriction struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reservation struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	StartDate   time.Time
	EndDate     time.Time
	RoomId      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Room        Room
}

type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomId        int
	ReservationId int
	RestrictionId int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}
