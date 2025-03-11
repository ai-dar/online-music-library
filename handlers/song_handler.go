package handlers

import (
	"net/http"
	"online-music-library/config"
	"online-music-library/repositories"
	"online-music-library/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllSongs получает список всех песен
// @Summary Получить список песен
// @Description Возвращает список всех доступных песен
// @Tags songs
// @Produce json
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]string
// @Router /songs [get]
func GetAllSongs(c *gin.Context) {
	songs, err := repositories.GetAllSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения песен"})
		return
	}
	c.JSON(http.StatusOK, songs)
}

// AddSong добавляет новую песню
// @Summary Добавить песню
// @Description Добавляет новую песню в базу данных
// @Tags songs
// @Accept json
// @Produce json
// @Param request body AddSongRequest true "Группа и название песни"
// @Success 201 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs [post]
func AddSong(c *gin.Context) {
	var req AddSongRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	cfg := config.LoadConfig()
	song, err := services.CreateSong(cfg, req.Group, req.Song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления песни"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

// AddSongRequest структура запроса для добавления песни
type AddSongRequest struct {
	Group string `json:"group" example:"Queen"`
	Song  string `json:"song" example:"Bohemian Rhapsody"`
}

// DeleteSong удаляет песню по ID
// @Summary Удалить песню
// @Description Удаляет песню по ID
// @Tags songs
// @Param id path int true "ID песни"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs/{id} [delete]
func DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repositories.DeleteSong(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Песня удалена"})
}
