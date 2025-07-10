package main

import (
	"music-queue/config"
	"music-queue/server"
)

func main() {
	config.LoadEnv()
	dbUrl := config.Getenv("DATABASE_URL")
	server.ConnectDB(dbUrl)

	server.StartGrpcServer()
}

// protoc --go-grpc_out=. --go_out=. *.proto
