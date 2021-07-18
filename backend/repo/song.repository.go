package repo

import (
	"log"

	"github.com/Zubayear/song-store/entity"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SongRepository interface {
	CreateSong(song *entity.Song) (*entity.Song, error)
	GetAllSongs() (*[]entity.Song, error)
	UpdateSong(song *entity.Song, songID uint) (*entity.Song, error)
	DeleteSongPermanently(songID uint) (string, error)
	GetSong(songID uint) (*entity.Song, error)
}

type database struct {
	dbConn *gorm.DB
}

func NewSongRepository() SongRepository {
	dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Exec("CREATE DATABASE IF NOT EXISTS " + "songdb").Exec("USE " + "songdb")
	if err != nil {
		panic("failed to connect to database")
	}
	err = db.AutoMigrate(&entity.Song{})
	if err != nil {
		log.Fatalf("db.AutoMigrate: %v", err)
		return nil
	}
	return &database{
		dbConn: db,
	}
}

func (d *database) CreateSong(song *entity.Song) (*entity.Song, error) {
	result := d.dbConn.Create(&song)
	if err := result.Error; err != nil {
		return nil, err
	}
	return song, nil
}

func (d *database) GetAllSongs() (*[]entity.Song, error) {
	var songs *[]entity.Song
	result := d.dbConn.Find(&songs)
	if err := result.Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (d *database) UpdateSong(song *entity.Song, songID uint) (*entity.Song, error) {
	var updatedSong *entity.Song
	result := d.dbConn.First(&updatedSong, songID)
	if err := result.Error; err != nil {
		return nil, err
	}
	if updatedSong.SongName == "" {
		updatedSong.SongName = song.SongName
	}
	if updatedSong.SongDuration == 0 {
		updatedSong.SongDuration = song.SongDuration
	}
	if updatedSong.SongHits == 0 {
		updatedSong.SongHits = song.SongHits
	}
	update := d.dbConn.Save(&updatedSong)
	if err := update.Error; err != nil {
		return nil, err
	}
	return updatedSong, nil
}

func (d *database) DeleteSongPermanently(songID uint) (string, error) {
	result := d.dbConn.Unscoped().Delete(&entity.Song{}, songID)
	if err := result.Error; err != nil {
		return "Sorry but couldn't delete this song", err
	}
	return "Song deleted permanently", nil
}

func (d *database) GetSong(songID uint) (*entity.Song, error) {
	var song *entity.Song
	result := d.dbConn.First(&song, songID)
	if err := result.Error; err != nil {
		return nil, err
	}
	return song, nil
}
