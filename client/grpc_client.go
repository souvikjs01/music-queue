package client

import (
	"log"
	proto "music-queue/protoc"
	"sync"

	"google.golang.org/grpc"
)

var (
	Client proto.MusicQueueServiceClient
	once   sync.Once
)

func InitGrpcClient() {
	once.Do(func() {
		conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure())

		if err != nil {
			log.Fatalln("could not connect to grpc server", err)
		}

		Client = proto.NewMusicQueueServiceClient(conn)
		log.Println("connected to grpc server!")
	})
}
