package mysql

import (
	"bookings/pkg/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const maxNoConnectionsErrorsBeforeReinitialization = 250

type StorageRepository struct {
	Pool *sql.DB
}

var defaultPrimaryKeys = []string{"id"}
var reservationColumns = []string{"room_id", "people_id", "start_date", "end_date"}
var peopleColumns = []string{"first_name", "last_name", "phone", "email"}
var roomColumns = []string{"name", "description"}

func NewStorageRepository(username, password, databaseName string) (*StorageRepository, error) {
	if username == "" {
		username = "root"
	}
	if databaseName == "" {
		databaseName = "bookings"
	}

	connectionStr := username + ":" + password + "@/" + databaseName
	fmt.Println("str: ", connectionStr)

	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		panic(err)
	}
	s := new(StorageRepository)
	s.Pool = db
	return s, nil
}

// SaveReservation
func (sR *StorageRepository) SaveReservation(ctx context.Context, res models.Reservation) (models.Reservation, error) {
	query := "INSERT INTO reservations (" + strings.Join(reservationColumns, ",") + ") VALUES ( ?, ?, ?, ? )"
	fmt.Println("query: ", query)
	stmtIns, err := sR.Pool.Prepare(query)
	if err != nil {
		return res, err
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	something, err := stmtIns.Exec(res.RoomId, res.PeopleId, res.StartDate, res.EndDate)
	fmt.Printf("\n save reservation: %v", res)
	fmt.Printf("\n save something: %v", something)
	if err != nil {
		return res, err
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
	return res, nil
}

// SaveRoom
func (sR *StorageRepository) SaveRoom(ctx context.Context, room models.Room) (models.Room, error) {
	query := "INSERT INTO rooms (" + strings.Join(roomColumns, ",") + ") VALUES ( ?, ?)"
	fmt.Println("query: ", query)
	stmtIns, err := sR.Pool.Prepare(query)
	if err != nil {
		return room, err
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	something, err := stmtIns.Exec(room.Name, room.Description)
	fmt.Printf("\n save room: %v", room)
	fmt.Printf("\n save something: %v", something)
	if err != nil {
		return room, err
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
	return room, nil
}

// SavePerson
func (sR *StorageRepository) SavePerson(ctx context.Context, person models.People) (models.People, error) {
	query := "INSERT INTO people (" + strings.Join(peopleColumns, ",") + ") VALUES ( ?, ?, ?, ? )"
	fmt.Println("query: ", query)
	stmtIns, err := sR.Pool.Prepare(query)
	if err != nil {
		return person, err
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	something, err := stmtIns.Exec(person.FirstName, person.LastName, person.Phone, person.Email)
	fmt.Printf("\n save person: %v", person)
	fmt.Printf("\n save something: %v", something)
	if err != nil {
		return person, err
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
	return person, nil
}

// FindReservationsForPerson
func (sR *StorageRepository) FindReservationsForPerson(ctx context.Context, person models.People, returnFirst bool) ([]models.Reservation, error) {
	var results []models.Reservation
	var reso models.Reservation
	var personId int
	if person.Id > 0 {
		personId = person.Id
	} else {
		resultingPerson, _ := sR.FindPerson(ctx, person, true)

		if len(resultingPerson) == 0 {
			return results, nil
		}
		personId = person.Id
	}

	query := "SELECT " + strings.Join(defaultPrimaryKeys, ", ") + ", " + strings.Join(reservationColumns, ", ") + " FROM reservations"
	query = query + " WHERE people_id=?"
	fmt.Println("\nquery: ", query)
	fmt.Printf("for personId: %v", personId)
	rows, err := sR.Pool.Query(query, personId)

    defer rows.Close()

    for rows.Next() {
		var startDate string
		var endDate string
		err := rows.Scan(&reso.Id, &reso.RoomId, &reso.PeopleId, &startDate, &endDate)
		timeFormat := "2006-01-02" // "Jan 2, 2006 at 3:04pm (MST)"
		reso.StartDate, _ = time.Parse(timeFormat, startDate)
		reso.EndDate, _ = time.Parse(timeFormat, endDate)
		fmt.Printf("\nreso: %v, startDate: %v, endDate: %v", reso, startDate, endDate)
        if err != nil {
            // log.Fatal(err)
		}
		results = append(results, reso)
    }

    if err = rows.Err(); err != nil {
		return results, err
	}

	return results, nil
}

// FindRoom
func (sR *StorageRepository) FindRoom(ctx context.Context, name string) (models.Room, error) {
	var result models.Room

	query := "SELECT " + strings.Join(defaultPrimaryKeys, ", ") + ", " + strings.Join(roomColumns, ", ") + " FROM rooms"
	query = query + " WHERE name=?"
	fmt.Println("\nquery: ", query, name)
	row := sR.Pool.QueryRow(query, name)
	err := row.Scan(&result.Id, &result.Name, &result.Description)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindPerson
func (sR *StorageRepository) FindPerson(ctx context.Context, person models.People, returnFirst bool) ([]models.People, error) {
	var results []models.People
	query := "SELECT " + strings.Join(defaultPrimaryKeys, ", ") + ", " + strings.Join(peopleColumns, ", ") + " FROM people"
	
	var filters = []string{person.FirstName, person.LastName, person.Phone, person.Email}
	var filterStrings []string
	filterCounter := 0
	for i, str := range filters {
		if str != "" {
			filterCounter += 1
			filters[i] = "  " + peopleColumns[i] + " = ?"
			filterStrings = append(filterStrings, str)
		}
	}
	// fmt.Printf("\nfind person filters updated: %v", filters)
	if filterCounter > 0 {
		// query = query + " WHERE " + strings.Join(filters, " AND ") // need to skip blank filters
		query = query + " WHERE "
		firstFilter := false
		for _, filterStr := range filters {
			if filterStr != "" {
				if firstFilter {
					query = query + " AND "	+ filterStr
				} else {
					firstFilter = true
					query = query + " "	+ filterStr
				}
			}
		}
	}
	// fmt.Println("\nresulting query: ", query,"\nfilterStrings: ", filterStrings,"\nfilterCounter: ", filterCounter)
	var err error
	var rows *sql.Rows
	if filterCounter == 1 {
		rows, err = sR.Pool.Query(query, filterStrings[0])
	} else if filterCounter == 2 {
		rows, err = sR.Pool.Query(query, filterStrings[0], filterStrings[1])
	} else if filterCounter == 3 {
		rows, err = sR.Pool.Query(query, filterStrings[0], filterStrings[1], filterStrings[2])
	} else if filterCounter == 4 {
		rows, err = sR.Pool.Query(query, filterStrings[0], filterStrings[1], filterStrings[2], filterStrings[3])
	} else {
		rows, err = sR.Pool.Query(query)
	}

    defer rows.Close()

    for rows.Next() {
		err := rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Phone, &person.Email)
        if err != nil {
            // log.Fatal(err)
		}
		results = append(results, person)
    }

    if err = rows.Err(); err != nil {
		return results, err
	}

	return results, nil
}
