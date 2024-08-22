package services

import (
	"encoding/json"
	"net/http"

	"github.com/inihikam/web-sti-api/models"
)

func GetBkAnnounce() ([]models.BkAnnounce, error) {
	resp, err := http.Get("https://bimbingan-karir.inihikam.my.id/api/pengumuman")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bkAnnounce []models.BkAnnounce
	if err := json.NewDecoder(resp.Body).Decode(&bkAnnounce); err != nil {
		return nil, err
	}

	return bkAnnounce, nil
}
