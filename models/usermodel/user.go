package usermodel

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"ID" gorm:"type:uuid;primaryKey"`
	FirstName   string    `json:"FirstName" gorm:"type:varchar;size:100;not null"`
	LastName    string    `json:"LastName" gorm:"type:varchar;size:100;not null"`
	Username    string    `json:"Username" gorm:"type:varchar;size:100;not null;unique"`
	Email       string    `json:"Email"  gorm:"type:varchar;size:100;not null;unique"`
	PhoneNumber string    `json:"PhoneNumber"  gorm:"type:varchar;size:100;not null;unique"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}
