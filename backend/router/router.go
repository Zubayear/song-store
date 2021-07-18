package router

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Zubayear/song-store/controller"
	"github.com/Zubayear/song-store/graph"
	"github.com/Zubayear/song-store/graph/generated"
	"github.com/Zubayear/song-store/repo"
	"github.com/Zubayear/song-store/service"
	"github.com/gin-gonic/gin"
)

var (
	songRepository = repo.NewSongRepository()
	songService    = service.New(songRepository)
	songController = controller.New(songService)
)

func graphqlHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Routing() {
	router := gin.Default()
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.GET("/songs", songController.GetAllSongs)
	router.GET("/songs/:id", songController.GetSong)
	router.POST("/songs", songController.CreateSong)
	router.PUT("/songs/:id", songController.UpdateSong)
	router.DELETE("/songs/:id", songController.DeleteSongPermanently)
	router.GET("/query")
	err := router.Run(":9000")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
