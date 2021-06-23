package entity

import "gorm.io/gorm"

type Song struct {
	gorm.Model   `json:"-"`
	ID           uint
	SongName     string  `json:"song_name"`
	SongDuration float32 `json:"song_duration"`
	SongHits     int64   `json:"song_hits"`
}
