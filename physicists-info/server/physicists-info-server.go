package server

import (
	"context"
	pb "cosmological-constant/pb"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"physicists-info/models"
	"time"
)

var physicists = []models.Physicist{
	models.NewPhysicist("Albert", "Einstein", "Germany", time.UnixMilli(-2865422008), time.UnixMilli(-464107200), ""),
	models.NewPhysicist("John", "von Neumann", "Austria-Hungary", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Niels", "Bohr", "Denmark", time.UnixMilli(-2658148408), time.UnixMilli(-224692800), ""),
	models.NewPhysicist("Hermann", "Minkowski", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Robert", "Oppenheimer", "The United States", time.UnixMilli(-2073134400), time.UnixMilli(-90513600), ""),
	models.NewPhysicist("Lawrence Maxwell", "Krauss", "The United States", time.UnixMilli(-492273600), time.Now(), ""),
	models.NewPhysicist("Leo", "Szilard", "Austria-Hungary", time.UnixMilli(-2268484800), time.UnixMilli(-176395200), ""),
	models.NewPhysicist("James", "Franck", "Germany", time.UnixMilli(-2724935608), time.UnixMilli(-177172800), ""),
	models.NewPhysicist("Hans", "Bethe", "Germany", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Richard", "Feynman", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Murrary", "Gell-Mann", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Steven", "Weinberg", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Nathan", "Rosen", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Boris Yakovlevich", "Podolsky", "Russian Empire", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Leon Max", "Lederman", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Vitaly", " Ginzburg", "Russian Empire", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Hilde", "Levi", "Germany", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Sheldon", "Glashow", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Robert", "Serber", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), ""),
	models.NewPhysicist("Joseph", "Rotblat", "Russian Empire", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
	models.NewPhysicist("Otto Robert", "Frisch", "Austria-Hungary", time.UnixMilli(-1918737600), time.UnixMilli(865153200), "ADJUST DATE"),
}

type PhysicistsInfoServer struct {
	pb.UnimplementedPhysicistsInfoServer
}

func (p PhysicistsInfoServer) GetAllPhysicist(ctx context.Context, request *pb.AllPhysicistsRequest) (*pb.AllPhysicists, error) {
	var returnedPhysicists = []*pb.Physicist{}
	for _, physicist := range physicists {
		returnedPhysicists = append(returnedPhysicists, &pb.Physicist{
			PhysicistId:    wrapperspb.String(physicist.PhysicistId.String()),
			FirstName:      physicist.FirstName,
			LastName:       physicist.LastName,
			CountryOfBirth: physicist.CountryOfBirth,
			DateOfBirth:    physicist.DateOfBirth.UnixMilli(),
			DateOfDeath:    physicist.DateOfDeath.UnixMilli(),
			Biography:      physicist.Biography,
		})
	}
	return &pb.AllPhysicists{Physicists: returnedPhysicists}, status.New(codes.OK, "").Err()
}

func (p PhysicistsInfoServer) GetPhysicistById(ctx context.Context, uuid *wrapperspb.StringValue) (*pb.Physicist, error) {
	for _, physicist := range physicists {
		if uuid.Value == physicist.PhysicistId.String() {
			return &pb.Physicist{
				PhysicistId:    wrapperspb.String(physicist.PhysicistId.String()),
				FirstName:      physicist.FirstName,
				LastName:       physicist.LastName,
				CountryOfBirth: physicist.CountryOfBirth,
				DateOfBirth:    physicist.DateOfBirth.UnixMilli(),
				DateOfDeath:    physicist.DateOfDeath.UnixMilli(),
				Biography:      physicist.Biography,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Physicist with the given id does not exist.", uuid)
}

func (p PhysicistsInfoServer) AddPhysicist(ctx context.Context, physicist *pb.Physicist) (*pb.Physicist, error) {
	newPhysicist := models.Physicist{
		PhysicistId:    uuid.FromStringOrNil(physicist.PhysicistId.String()),
		FirstName:      physicist.FirstName,
		LastName:       physicist.LastName,
		CountryOfBirth: physicist.CountryOfBirth,
		DateOfBirth:    time.UnixMilli(physicist.DateOfBirth),
		DateOfDeath:    time.UnixMilli(physicist.DateOfDeath),
		Biography:      physicist.Biography,
	}

	physicists = append(physicists, newPhysicist)

	return physicist, status.New(codes.OK, "").Err()
}

func (p PhysicistsInfoServer) GetPhysicistsByCountryOfBirth(value *wrapperspb.StringValue, server pb.PhysicistsInfo_GetPhysicistsByCountryOfBirthServer) error {
	log.Println("The country's name we will be filtering by is ", value.Value)
	for _, physicist := range physicists {
		if physicist.CountryOfBirth == value.Value {
			err := server.Send(&pb.Physicist{
				PhysicistId:    wrapperspb.String(physicist.PhysicistId.String()),
				FirstName:      physicist.FirstName,
				LastName:       physicist.LastName,
				CountryOfBirth: physicist.CountryOfBirth,
				DateOfBirth:    physicist.DateOfBirth.UnixMilli(),
				DateOfDeath:    physicist.DateOfDeath.UnixMilli(),
				Biography:      physicist.Biography,
			})
			if err != nil {
				return fmt.Errorf("Error sending physicist to stream: %v", err)
			}
		}
	}
	return nil
}
