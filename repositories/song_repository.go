package repositories

import (
	"online-music-library/database"
	"online-music-library/models"
)

func GetAllSongs() ([]models.Song, error) {
	var songs []models.Song
	err := database.DB.Find(&songs).Error
	return songs, err
}

func GetSongByID(id uint) (models.Song, error) {
	var song models.Song
	err := database.DB.First(&song, id).Error
	return song, err
}

func CreateSong(song *models.Song) error {
	return database.DB.Create(song).Error
}

func DeleteSong(id uint) error {
	return database.DB.Delete(&models.Song{}, id).Error
}
