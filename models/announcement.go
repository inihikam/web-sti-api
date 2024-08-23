package models

import "time"

type Announcement struct {
	ID           uint   `json:"id"`
	Judul        string `json:"judul"`
	Penulis      string `json:"penulis"`
	Konten       string `json:"konten"`
	Published_at time.Time `json:"published_at"`
}