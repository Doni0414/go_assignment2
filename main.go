package main

import (
	"exercise1/db"
	"fmt"
	"os"
)

func main() {
	host := "localhost"
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := "go_db_exercise1"
	port := "5432"
	sslmode := "disable"
	timeZone := "Asia/Almaty"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timeZone=%s", host, user, password, dbname, port, sslmode, timeZone)

	db.Connect(dsn)

	instructors := []db.Instructor{
		{
			FullName: "Nurbol Sabitov",
			Age:      24,
		},
		{
			FullName: "Azamat Serek",
			Age:      28,
		},
	}
	for _, instructor := range instructors {
		db.CreateInstructor(&instructor)
	}
	ins1 := db.FindInstructorById(1)
	db.DeleteInstructor(&ins1)

	ins2 := db.FindInstructorById(2)
	instructorFieldsToBeUpdated := db.Instructor{
		Age: 34,
	}
	db.UpdateInstructor(&ins2, &instructorFieldsToBeUpdated)
}
