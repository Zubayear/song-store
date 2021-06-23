package router

import (
	"github.com/Zubayear/song-store/controller"
	"github.com/Zubayear/song-store/repo"
	"github.com/Zubayear/song-store/service"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	songRepository = repo.NewSongRepository()
	songService    = service.New(songRepository)
	songController = controller.New(songService)
)

func Routing() {
	router := gin.Default()
	router.GET("/songs", songController.GetAllSongs)
	router.GET("/songs/:id", songController.GetSong)
	router.POST("/songs", songController.CreateSong)
	router.PUT("/songs/:id", songController.UpdateSong)
	router.DELETE("/songs/:id", songController.DeleteSongPermanently)
	err := router.Run(":9000")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
