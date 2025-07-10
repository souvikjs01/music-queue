package models

type Song struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Artist     string
	Upvotes    int
	YoutubeURL string `gorm:"column:youtube_url"`
}
