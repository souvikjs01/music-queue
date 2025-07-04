package main

import (
	"log"
	"music-queue/config"
	"music-queue/server"
	"net"

	proto "music-queue/protoc"

	"google.golang.org/grpc"
)

func main() {
	config.LoadEnv()
	dbUrl := config.Getenv("DATABASE_URL")
	server.ConnectDB(dbUrl)

	// start grpc server
	lis, err := net.Listen("tcp", ":8091")
	if err != nil {
		log.Fatalln("Failed to listen")
	}

	grpcServer := grpc.NewServer()
	musicServer := server.NewMusicQueueServiceServer()

	proto.RegisterMusicQueueServiceServer(grpcServer, musicServer)

	log.Println("gRPC server running on the port 8091")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}

}
