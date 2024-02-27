package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(dsn string) {
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func MigrateAllTables() {
	db.AutoMigrate(&Student{}, &Course{}, &Department{}, &Enrollment{}, &Instructor{})
}

func MigrateTable(table interface{}) {
	db.AutoMigrate(&table)
}
