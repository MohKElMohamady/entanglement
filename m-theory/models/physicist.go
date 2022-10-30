package models

import (
	//uuid "github.com/satori/go.uuid"
	"time"
)

type Physicist struct {
	//PhysicistId    uuid.UUID
	FirstName      string
	LastName       string
	CountryOfBirth string
	DateOfBirth    time.Time
	DateOfDeath    time.Time
	Biography      string
}

func NewPhysicist(firstName string, lastName string, countryOfBirth string, dateOfBirth time.Time, dateOfDeath time.Time, biography string) Physicist {
	return Physicist{
		//PhysicistId:    uuid.UUID{},
		FirstName:      firstName,
		LastName:       lastName,
		CountryOfBirth: countryOfBirth,
		DateOfBirth:    time.Time{},
		DateOfDeath:    time.Time{},
		Biography:      biography,
	}
}
