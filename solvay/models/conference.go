package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Europe map[string]bool

type Conference struct {
	ConferenceId            uuid.UUID
	Name                    string
	Location                string
	Description             string
	StartTime               time.Time
	EndTime                 time.Time
	ParticipatingPhysicists []*participatingPhysicist
}

func NewConference(name string, location string, description string, startTime time.Time, endTime time.Time) *Conference {
	return &Conference{
		ConferenceId:            uuid.NewV1(),
		Name:                    name,
		Location:                location,
		Description:             description,
		StartTime:               startTime,
		EndTime:                 endTime,
		ParticipatingPhysicists: []*participatingPhysicist{},
	}
}

type participatingPhysicist struct {
	ParticipatingPhysicistId uuid.UUID
	FirstName                string
	LastName                 string
	CountryOfBirth           string
}

func NewParticipatingPhysicist(participatingPhysicistsID uuid.UUID, firstName string, lastName string, countryOfBirth string) *participatingPhysicist {
	return &participatingPhysicist{
		ParticipatingPhysicistId: participatingPhysicistsID,
		FirstName:                firstName,
		LastName:                 lastName,
		CountryOfBirth:           countryOfBirth,
	}
}
