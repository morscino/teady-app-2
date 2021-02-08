package postgres

import (
	"github.com/morscino/teady-2-app/utility/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(database config.Database) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=" + database.User + " password=" + database.Password + " dbname=" + database.Name + " sslmode=" + database.SSLMode,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
