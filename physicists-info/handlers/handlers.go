package handlers

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func RequestLoggerHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Server interceptor " + info.FullMethod)
	//m, err := handler(ctx, req)
	log.Println("Post Proc Message: %s")
	return nil, nil
}
