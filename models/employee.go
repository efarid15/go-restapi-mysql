package models

import (
	"time"
)

type Employee struct {
	ID        int       `json:"id"`
	IDNumber  string       `json:"id_number"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
