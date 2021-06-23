package repo

import (
	"github.com/Zubayear/song-store/entity"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	// db, err := gorm.Open("mysql", "root@tcp(127.0.0.1)/song_db?charset=utf8")
	dsn := "root:@tcp(127.0.0.1:3306)/songdb?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = db.AutoMigrate(&entity.Song{})
	if err != nil {
		log.Fatalln(err)
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
	updatedSong.SongName = song.SongName
	updatedSong.SongDuration = song.SongDuration
	updatedSong.SongHits = song.SongHits
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
