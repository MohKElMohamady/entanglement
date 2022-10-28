package server

import (
	context "context"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"physicists-info/models"
	pb "physicists-info/pb"

	"time"
)

var physicists = []models.Physicist{
	models.NewPhysicist("Albert", "Einstein", "Germany", time.UnixMilli(-2865422008), time.UnixMilli(-464107200), ""),
	models.NewPhysicist("Niels", "Bohr", "Denmark", time.UnixMilli(-2658148408), time.UnixMilli(-224692800), ""),
	models.NewPhysicist("Robert", "Oppenheimer", "The United States", time.UnixMilli(-2073134400), time.UnixMilli(-90513600), ""),
	models.NewPhysicist("Lawrence Maxwell", "Krauss", "The United States", time.UnixMilli(-492273600), time.Now(), ""),
	models.NewPhysicist("Leo", "Szilard", "Hungary", time.UnixMilli(-2268484800), time.UnixMilli(-176395200), ""),
	models.NewPhysicist("James", "Franck", "Germany", time.UnixMilli(-2724935608), time.UnixMilli(-177172800), ""),
	models.NewPhysicist("Robert", "Serber", "The United States", time.UnixMilli(-1918737600), time.UnixMilli(865153200), ""),
}

type PhysicistsInfoServer struct {
	pb.UnimplementedPhysicistsInfoServer
}

func (p PhysicistsInfoServer) GetAllPhysicist(ctx context.Context, request *pb.AllPhysicistsRequest) (*pb.AllPhysicists, error) {
	var returnedPhysicists = []*pb.Physicist{}
	for _, physicist := range physicists {
		returnedPhysicists = append(returnedPhysicists, &pb.Physicist{
			PhysicistId:    &pb.UUID{UuidInString: physicist.PhysicistId.String()},
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

func (p PhysicistsInfoServer) GetPhysicistById(ctx context.Context, uuid *pb.UUID) (*pb.Physicist, error) {
	for _, physicist := range physicists {
		if uuid.UuidInString == physicist.PhysicistId.String() {
			return &pb.Physicist{
				PhysicistId:    &pb.UUID{UuidInString: physicist.PhysicistId.String()},
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
		PhysicistId:    uuid.FromStringOrNil(physicist.PhysicistId.UuidInString),
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
