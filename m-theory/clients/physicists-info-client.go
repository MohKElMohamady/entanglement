package clients

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"m-theory/pb"
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
	return c.client.GetPhysicistById(ctx, &pb.UUID{UuidInString: uuid})
}

func (c physicistsInfoClient) GetAllPhysicists() (*pb.AllPhysicists, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.client.GetAllPhysicist(ctx, &pb.AllPhysicistsRequest{})
}
