package server

import (
	"context"
	"music-queue/models"
	proto "music-queue/protoc"
)

type MusicQueueServiceServer struct {
	proto.UnimplementedMusicQueueServiceServer
}

func NewMusicQueueServiceServer() *MusicQueueServiceServer {
	return &MusicQueueServiceServer{}
}

func (s *MusicQueueServiceServer) AddSong(ctx context.Context, req *proto.AddSongRequest) (*proto.AddSongResponse, error) {
	song := models.Song{
		Title:      req.GetTitle(),
		Artist:     req.GetArtist(),
		Upvotes:    0,
		YoutubeURL: req.YoutubeUrl,
	}
	if err := DB.Create(&song).Error; err != nil {
		return nil, err
	}

	return &proto.AddSongResponse{Message: "Song added successfully"}, nil
}

func (s *MusicQueueServiceServer) GetQueue(ctx context.Context, req *proto.Empty) (*proto.SongQueueResponse, error) {
	var songs []models.Song

	if err := DB.Order("upvotes DESC").Find(&songs).Error; err != nil {
		return nil, err
	}

	var protoSongs []*proto.Song

	for _, song := range songs {
		protoSongs = append(protoSongs, &proto.Song{
			Id:         int32(song.ID),
			Title:      song.Title,
			Artist:     song.Artist,
			Upvotes:    int32(song.Upvotes),
			YoutubeUrl: song.YoutubeURL,
		})
	}

	return &proto.SongQueueResponse{
		Songs: protoSongs,
	}, nil
}

func (s *MusicQueueServiceServer) UpvoteSong(ctx context.Context, req *proto.UpvoteRequest) (*proto.UpvoteResponse, error) {
	var song models.Song
	if err := DB.First(&song, req.GetSongId()).Error; err != nil {
		return nil, err
	}

	song.Upvotes += 1
	if err := DB.Save(&song).Error; err != nil {
		return nil, err
	}

	return &proto.UpvoteResponse{
		Message: "Upvoted successfully",
	}, nil
}

func (s *MusicQueueServiceServer) DeleteSong(ctx context.Context, req *proto.DeleteSongRequest) (*proto.DeleteSongResponse, error) {
	songId := req.GetSongId()

	var song models.Song
	if err := DB.Delete(&song, songId).Error; err != nil {
		return nil, err
	}

	return &proto.DeleteSongResponse{
		Message: "Song deleted successfully",
	}, nil

}
