package routes

import (
	"online-music-library/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/songs", handlers.GetAllSongs)
	r.POST("/songs", handlers.AddSong)
	r.DELETE("/songs/:id", handlers.DeleteSong)

	return r
}
