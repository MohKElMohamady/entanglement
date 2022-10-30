package clients

import (
	"context"
	"cosmological-constant/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"time"
)

const (
	physicistInfoServiceAddress = "localhost:50051"
)

var PhysicistsInfoClient = *NewPhysicistsInfoClient()

type physicistsInfoClient struct {
	client pb.PhysicistsInfoClient
}

func NewPhysicistsInfoClient() *physicistsInfoClient {
	connection, err := grpc.Dial(physicistInfoServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Failed to connect to PhysicistsInfoServer: %v", err.Error())
	}

	return &physicistsInfoClient{
		client: pb.NewPhysicistsInfoClient(connection),
	}
}

func (c physicistsInfoClient) GetPhysicistByUuid(uuid string) (*pb.Physicist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.client.GetPhysicistById(ctx, &wrapperspb.StringValue{Value: uuid})
}

func (c physicistsInfoClient) GetAllPhysicists() (*pb.AllPhysicists, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.client.GetAllPhysicist(ctx, &pb.AllPhysicistsRequest{})
}

func (c physicistsInfoClient) GetPhysicistByCountryOfBirth(countryOfBirth string) []pb.Physicist {
	var physcistsOfThatCountry = []pb.Physicist{}
	ctx := context.Background()
	physicistStream, _ := c.client.GetPhysicistsByCountryOfBirth(ctx, &wrapperspb.StringValue{Value: countryOfBirth})

	for {
		physicist, err := physicistStream.Recv()
		if err == io.EOF {
			break
		}
		physcistsOfThatCountry = append(physcistsOfThatCountry, *physicist)
	}
	for _, physicist := range physcistsOfThatCountry {
		log.Println("Found physicists who are named " + physicist.FirstName + " " + physicist.LastName + " living in " + countryOfBirth)
	}
	return physcistsOfThatCountry
}
