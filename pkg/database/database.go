package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	//syntax to connect to the database
	url := "postgres://postgres:sherlock16@localhost:5432/swiggy_model"
	connection, err := gorm.Open(postgres.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db = connection
	return db, err
}

func GetDB() *gorm.DB {
	return db
}
