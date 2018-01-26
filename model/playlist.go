package model

type Playlist struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	VideosCount int    `json:"videos_count"`
}
