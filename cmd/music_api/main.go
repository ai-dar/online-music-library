package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongResponse struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func main() {
	r := gin.Default()

	r.GET("/info", func(c *gin.Context) {
		group := c.Query("group")
		song := c.Query("song")

		if group == "" || song == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing group or song parameters"})
			return
		}

		// Заглушка (можно заменить на логику получения данных)
		response := SongResponse{
			Group:       group,
			Song:        song,
			ReleaseDate: "1968",
			Text:        "Hey Jude, don't make it bad...",
			Link:        "https://example.com/hey-jude",
		}

		c.JSON(http.StatusOK, response)
	})

	r.Run(":3000") // Запускаем сервер на порту 3000
}
