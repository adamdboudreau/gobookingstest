package models

type Reservation struct {
	Id        int
	RoomId    int
	PeopleId  int
	StartDate time.Time
	EndDate   time.Time
}
