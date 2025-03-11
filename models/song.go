package models

import (
	"time"
)

type Song struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate string    `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
