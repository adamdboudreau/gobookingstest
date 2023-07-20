package mysql

import (
	"bookings/pkg/models"
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const maxNoConnectionsErrorsBeforeReinitialization = 250

type StorageRepository struct {
	Pool *sql.DB
}

var reservationColumns = []string{"room_id", "last_name", "first_name", "email", "phone"}
var reservationPrimaryKeys = []string{"id"}

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

//load reservation
func (sR *StorageRepository) FindReservation(ctx context.Context, email string) (models.Reservation, error) {
	var res models.Reservation
	query := "SELECT " + strings.Join(reservationPrimaryKeys, ", ") + ", " + strings.Join(reservationColumns, ", ")
	query = query + " FROM reservations WHERE email = ?"
	stmtOut, err := sR.Pool.Prepare(query)
	fmt.Println("query: ", query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(email).Scan(&res.Id, &res.RoomId, &res.LastName, &res.FirstName, &res.Email, &res.Phone)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The reservation room id: %d, fn: %s, ln: %s", res.RoomId, res.FirstName, res.LastName)
	return res, nil
}

// save reservation
func (sR *StorageRepository) SaveReservation(ctx context.Context, res models.Reservation) {
	// cql := "UPDATE reservations SET " + strings.Join(reservationColumns, "=?, ") + "=? WHERE " + strings.Join(reservationPrimaryKeys, "=? AND ") + "=?"
	query := "INSERT INTO reservations (" + strings.Join(reservationColumns, ",") + ") VALUES ( ?, ?, ?, ?, ? )"
	fmt.Println("query: ", query)
	stmtIns, err := sR.Pool.Prepare(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(res.RoomId, res.LastName, res.FirstName, res.Email, res.Phone) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
