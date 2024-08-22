package models

import "time"

type AlumniAnnounce struct {
	Title     string    `json:"Posisi"`
	User      string    `json:"NamaPerusahaan"`
	Content   string    `json:"Deskripsi"`
	Date 	  time.Time `json:"created_at"`
}