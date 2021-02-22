package usermodel

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"ID" gorm:"type:uuid;primaryKey"`
	FirstName   string    `json:"FirstName" gorm:"type:varchar(100);not null"`
	LastName    string    `json:"LastName" gorm:"type:varchar(100);not null"`
	Username    string    `json:"Username" gorm:"type:varchar(100);not null;unique"`
	Email       string    `json:"Email"  gorm:"type:varchar(100);not null;unique"`
	PhoneNumber string    `json:"PhoneNumber"  gorm:"type:varchar(100);unique"`
	Password    []byte    `json:"Password" gorm:"type:bytea;not null"`
	Salt        []byte    `json:"Salt" gorm:"type:bytea;not null"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}
