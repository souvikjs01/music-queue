package client

import (
	"context"
	proto "music-queue/protoc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddSong(c *gin.Context) {
	var req struct {
		Title  string `json:"title"`
		Artist string `json:"artist"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	grpcReq := &proto.AddSongRequest{
		Title:  req.Title,
		Artist: req.Artist,
	}

	res, err := Client.AddSong(context.Background(), grpcReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.GetMessage(),
	})
}

func GetQueue(c *gin.Context) {
	res, err := Client.GetQueue(context.Background(), &proto.Empty{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res.GetSongs())
}

func UpvoteSong(c *gin.Context) {
	idParam := c.Param("id")
	songID, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	grpcReq := &proto.UpvoteRequest{
		SongId: int32(songID),
	}

	res, err := Client.UpvoteSong(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.GetMessage(),
	})
}
