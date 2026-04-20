package model

import (
	"time"

	"github.com/google/uuid"
)


type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string
	Email     string    `gorm:"unique"`
	Password  string 	`json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}