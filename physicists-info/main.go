package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"physicists-info/pb"
	s "physicists-info/server"
)

const (
	port = ":50051"
)

func main() {
	listener, error := net.Listen("tcp", port)
	if error != nil {
		log.Fatalln("Could not listen to the port, maybe it is already used?")
	}

	server := grpc.NewServer()
	pb.RegisterPhysicistsInfoServer(server, &s.PhysicistsInfoServer{})

	log.Println("Starting gRPC listener on port", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalln("Failed to serve the server: %v", err.Error())
	}
}
