package service

import (
	"github.com/Zubayear/song-store/entity"
	"github.com/Zubayear/song-store/repo"
)

type SongService interface {
	CreateSong(song *entity.Song) (*entity.Song, error)
	GetAllSongs() (*[]entity.Song, error)
	UpdateSong(song *entity.Song, songID uint) (*entity.Song, error)
	DeleteSongPermanently(songID uint) (string, error)
	GetSong(songID uint) (*entity.Song, error)
}

type songService struct {
	songRepository repo.SongRepository
}

func (ss *songService) CreateSong(song *entity.Song) (*entity.Song, error) {
	return ss.songRepository.CreateSong(song)
}

func (ss *songService) GetAllSongs() (*[]entity.Song, error) {
	return ss.songRepository.GetAllSongs()
}

func (ss *songService) UpdateSong(song *entity.Song, songID uint) (*entity.Song, error) {
	return ss.songRepository.UpdateSong(song, songID)
}

func (ss *songService) DeleteSongPermanently(songID uint) (string, error) {
	return ss.songRepository.DeleteSongPermanently(songID)
}

func (ss *songService) GetSong(songID uint) (*entity.Song, error) {
	return ss.songRepository.GetSong(songID)
}

func New(repo repo.SongRepository) SongService {
	return &songService{
		songRepository: repo,
	}
}
