package main

import (
	"cosmological-constant/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"solvay/server"
)

const (
	port = ":50052"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Could not listen to the port, maybe it is already used?")
	}

	solvay := grpc.NewServer()

	pb.RegisterSolvayServiceServer(solvay, &server.SolvayServer{})
	log.Println("Starting gRPC listener on port", port)
	if err := solvay.Serve(listener); err != nil {
		log.Fatalln("Failed to serve the server: %v", err.Error())
	}
}
