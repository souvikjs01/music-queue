package main

import (
	"log"
	"music-queue/api"
	"music-queue/client"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize gRPC Client
	client.InitGrpcClient()

	router := gin.Default()

	// Admin Routes
	router.POST("/admin/songs", api.AddSong)

	// User Routes
	router.GET("/songs", api.GetQueue)
	router.POST("/songs/:id/upvote", api.UpvoteSong)

	log.Println("API Gateway running on port 8080")
	router.Run(":8080")
}
