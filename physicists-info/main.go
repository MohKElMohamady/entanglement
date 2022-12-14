package main

import (
	"cosmological-constant/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"physicists-info/handlers"
	"physicists-info/server"
)

const (
	port = ":50051"
)

func main() {
	listener, error := net.Listen("tcp", port)
	if error != nil {
		log.Fatalln("Could not listen to the port, maybe it is already used?")
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(handlers.RequestLoggerHandler))
	pb.RegisterPhysicistsInfoServer(s, server.PhysicistsInfoServer{})
	log.Println("Starting gRPC listener on port", port)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve the server: %v", err.Error())
	}
}
