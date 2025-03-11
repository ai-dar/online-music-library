package services

import (
	"log"
	"online-music-library/config"
	"online-music-library/models"
	"online-music-library/repositories"

	"github.com/go-resty/resty/v2"
)

func FetchSongDetails(cfg config.Config, group, song string) (models.Song, error) {
	client := resty.New()
	var songDetail models.Song

	resp, err := client.R().
		SetQueryParams(map[string]string{"group": group, "song": song}).
		SetResult(&songDetail).
		Get(cfg.MusicAPI)

	if err != nil {
		log.Println("Ошибка при запросе к API:", err)
		return models.Song{}, err
	}

	log.Println("Ответ от API:", resp)
	return songDetail, nil
}

func CreateSong(cfg config.Config, group, song string) (models.Song, error) {
	songDetail, err := FetchSongDetails(cfg, group, song)
	if err != nil {
		return models.Song{}, err
	}

	newSong := models.Song{
		Group:       group,
		Song:        song,
		ReleaseDate: songDetail.ReleaseDate,
		Text:        songDetail.Text,
		Link:        songDetail.Link,
	}

	err = repositories.CreateSong(&newSong)
	if err != nil {
		log.Println("Ошибка сохранения песни:", err)
	}
	return newSong, err
}
