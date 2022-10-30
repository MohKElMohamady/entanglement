package server

import (
	"cosmological-constant/pb"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"solvay/models"
	"time"
)

type SolvayServer struct {
	pb.UnimplementedSolvayServiceServer
}

var europe = models.Europe{
	"Germany":         true,
	"Austria-Hungary": true,
	"Denmark":         true,
	"Russian Empire":  true,
}

var solvayConference = models.NewConference("Solvay Conference", "Belgium", "", time.UnixMilli(-1835793600), time.Now())

var siliconValleyAstronomyLectures = models.NewConference("Silicon Valley Astronomy Lectures", "California", "", time.UnixMilli(-1835793600), time.Now())

var manhattanProject = models.NewConference("Manhattan Project", "The United States", "", time.UnixMilli(-864229200), time.UnixMilli(-706290000))

var manhattan = pb.Conference{
	ConferenceId:           wrapperspb.String(manhattanProject.ConferenceId.String()),
	Name:                   manhattanProject.Name,
	Location:               manhattanProject.Location,
	Description:            manhattanProject.Description,
	StartTime:              manhattanProject.StartTime.UnixMilli(),
	EndTime:                manhattanProject.EndTime.UnixMilli(),
	PhysicistsInConference: []*pb.Physicist{},
}

var solvay = pb.Conference{
	ConferenceId:           wrapperspb.String(solvayConference.ConferenceId.String()),
	Name:                   manhattanProject.Name,
	Location:               manhattanProject.Location,
	Description:            manhattanProject.Description,
	StartTime:              manhattanProject.StartTime.UnixMilli(),
	EndTime:                manhattanProject.EndTime.UnixMilli(),
	PhysicistsInConference: []*pb.Physicist{},
}

func (s SolvayServer) GetConferencesForPhysicists(server pb.SolvayService_GetConferencesForPhysicistsServer) error {
	var designatedConferences *pb.DesignatedConferences

	for {
		physicist, err := server.Recv()
		if err == io.EOF {
			for _, participatingPhysicist := range manhattanProject.ParticipatingPhysicists {
				manhattan.PhysicistsInConference = append(manhattan.PhysicistsInConference, &pb.Physicist{
					PhysicistId:    wrapperspb.String(participatingPhysicist.ParticipatingPhysicistId.String()),
					FirstName:      participatingPhysicist.FirstName,
					LastName:       participatingPhysicist.LastName,
					CountryOfBirth: participatingPhysicist.CountryOfBirth,
				})
			}
			for _, participatingPhysicist := range solvayConference.ParticipatingPhysicists {
				manhattan.PhysicistsInConference = append(manhattan.PhysicistsInConference, &pb.Physicist{
					PhysicistId:    wrapperspb.String(participatingPhysicist.ParticipatingPhysicistId.String()),
					FirstName:      participatingPhysicist.FirstName,
					LastName:       participatingPhysicist.LastName,
					CountryOfBirth: participatingPhysicist.CountryOfBirth,
				})
			}
			designatedConferences.Conferences = append(designatedConferences.Conferences, &manhattan)
			designatedConferences.Conferences = append(designatedConferences.Conferences, &solvay)
			designatedConferences.Conferences = append(designatedConferences.Conferences)
			server.SendAndClose(designatedConferences)
		}
		if europe[physicist.CountryOfBirth] {
			solvayConference.ParticipatingPhysicists = append(manhattanProject.ParticipatingPhysicists, models.NewParticipatingPhysicist(uuid.FromStringOrNil(physicist.PhysicistId.GetValue()), physicist.FirstName, physicist.LastName, physicist.GetCountryOfBirth()))
		} else {
			manhattanProject.ParticipatingPhysicists = append(manhattanProject.ParticipatingPhysicists, models.NewParticipatingPhysicist(uuid.FromStringOrNil(physicist.PhysicistId.GetValue()), physicist.FirstName, physicist.LastName, physicist.GetCountryOfBirth()))
		}
	}
	return nil
}
