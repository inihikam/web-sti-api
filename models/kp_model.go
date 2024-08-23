package models

type KpAnnounce struct {
	ID      uint   `json:"id"`
	Title   string `json:"judul"`
	User    string `json:"user"`
	Content string `json:"isi"`
	Date    string `json:"published_at"`
}