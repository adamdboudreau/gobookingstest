package main

import (
	"time"
	"fmt"
	"context"
	// "bookings/pkg/config"
	// "bookings/pkg/forms"
	// "bookings/pkg/helpers"
	"bookings/pkg/models"
	// "bookings/pkg/render"
	"bookings/pkg/storage/mysql"
)
// go run cmd/scripts/seed/main.go
func main() {
	fmt.Println("start seeding default data into db")

	mysqlDB, err := mysql.NewStorageRepository("", "", "")
	if err != nil {
		// log.Fatal("cannot load templates")
		return// err
	}
	ctx := context.Background()

	testRoom := models.Room{Name: "Majors Room", Description: "Majors Room Desc"}
	fmt.Printf("save Record: %v", testRoom)
	resRoom, err := mysqlDB.SaveRoom(ctx, testRoom)
	fmt.Printf("resRoom: %v\nerr: %v", resRoom, err)

	testPerson := models.People{FirstName: "Adam", LastName: "Smith", Email: "adam@demo.com", Phone: "123455"}
	fmt.Printf("save Record: %v", testPerson)
	resPerson, err := mysqlDB.SavePerson(ctx, testPerson)
	fmt.Printf("resPerson: %v\nerr: %v", resPerson, err)

	testFindPeople(mysqlDB, testPerson)

	startDate := time.Now()
	endDate := startDate.Add(72 * time.Hour)
	testReservation := models.Reservation{RoomId: 1, PeopleId: 1, StartDate: startDate, EndDate: endDate}
	fmt.Printf("save Record: %v", testReservation)
	resReservation, err := mysqlDB.SaveReservation(ctx, testReservation)
	fmt.Printf("resReservation: %v\nerr: %v", resReservation, err)

}

func testFindPeople(mysqlDB *mysql.StorageRepository, testPerson models.People) {
	ctx := context.Background()
	people, err := mysqlDB.FindPerson(ctx, testPerson, false)
	fmt.Printf("find person err: %v", err)
	for _, person := range people {
		fmt.Printf("4q found person: %v", person)
	}

	testPerson.Id = 1
	reservations, err := mysqlDB.FindReservationsForPerson(ctx, testPerson, false)
	fmt.Printf("\nfind reservations err: %v", err)
	for _, reso := range reservations {
		fmt.Printf("reso: %v", reso)
	}

	room, err := mysqlDB.FindRoom(ctx, "Majors Room")
	fmt.Printf("\nfind room err: %v", err)
	fmt.Printf("room: %v", room)

	// testPerson = models.People{FirstName: "Adam", LastName: "Boudreau", Email: "adam@demo.com", Phone: ""}
	// people, err = mysqlDB.FindPerson(ctx, testPerson, false)
	// fmt.Printf("find person err: %v", err)
	// for _, person := range people {
	// 	fmt.Printf("3q found person: %v", person)
	// }

	// testPerson = models.People{FirstName: "Adam", LastName: "Boudreau", Email: "", Phone: ""}
	// people, err = mysqlDB.FindPerson(ctx, testPerson, false)
	// fmt.Printf("find person err: %v", err)
	// for _, person := range people {
	// 	fmt.Printf("2q found person: %v", person)
	// }

	// testPerson = models.People{FirstName: "Adam", LastName: "", Email: "", Phone: ""}
	// people, err = mysqlDB.FindPerson(ctx, testPerson, false)
	// fmt.Printf("find person err: %v", err)
	// for _, person := range people {
	// 	fmt.Printf("1q found person: %v", person)
	// }
}