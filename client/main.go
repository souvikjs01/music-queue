package main

import (
	"log"
	client "music-queue/client/grpc_client"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize gRPC Client
	client.InitGrpcClient()

	router := gin.Default()

	router.Use(cors.Default())

	// Admin Routes
	router.POST("/admin/songs", client.AddSong)

	// User Routes
	router.GET("/songs", client.GetQueue)
	router.POST("/songs/:id/upvote", client.UpvoteSong)
	router.DELETE("/songs/remove/:id", client.DeleteSong)

	log.Println("API Gateway running on port 8080")
	router.Run(":8080")
}
