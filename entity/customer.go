package entity

import "time"

type Customer struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
