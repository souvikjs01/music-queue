package server

import (
	"log"
	proto "music-queue/protoc"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcServer() {
	listener, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	musicService := NewMusicQueueServiceServer()

	proto.RegisterMusicQueueServiceServer(grpcServer, musicService)
	log.Println("gRPC server is running on the port 8091")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
