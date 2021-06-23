package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Zubayear/song-store/entity"
	"github.com/Zubayear/song-store/service"
	"github.com/gin-gonic/gin"
)

type SongController interface {
	CreateSong(c *gin.Context)
	GetAllSongs(c *gin.Context)
	GetSong(c *gin.Context)
	UpdateSong(c *gin.Context)
	DeleteSongPermanently(c *gin.Context)
}

type controller struct {
	service service.SongService
}

func (con *controller) GetAllSongs(c *gin.Context) {
	songs, err := con.service.GetAllSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalln(err)
		return
	}
	data, err := json.Marshal(songs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatalln(err)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(data)
}

func (con *controller) CreateSong(c *gin.Context) {
	song, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var songToAdd *entity.Song
	err = json.Unmarshal(song, &songToAdd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := con.service.CreateSong(songToAdd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := json.Marshal(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(data)
}

func (con *controller) GetSong(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := con.service.GetSong(uint(songID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := json.Marshal(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(data)
}

func (con *controller) UpdateSong(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	uSong, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var updatedSong *entity.Song
	err = json.Unmarshal(uSong, &updatedSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := con.service.UpdateSong(updatedSong, uint(songID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := json.Marshal(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(data)
}

func (con *controller) DeleteSongPermanently(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := con.service.DeleteSongPermanently(uint(songID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}

func New(songService service.SongService) SongController {
	return &controller{
		service: songService,
	}
}
