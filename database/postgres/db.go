package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=tams password=tams dbname=teady_db sslmode=disable ",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
