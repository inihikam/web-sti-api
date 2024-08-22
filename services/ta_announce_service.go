package services

import (
	"encoding/json"
	"net/http"
	"github.com/inihikam/web-sti-api/models"
)

func GetTaAnnounce() ([]models.TaAnnounce, error) {
	resp, err := http.Get("https://bimbingan-online.inihikam.my.id/api/pengumuman")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var taAnnounce []models.TaAnnounce
	if err := json.NewDecoder(resp.Body).Decode(&taAnnounce); err != nil {
		return nil, err
	}

	return taAnnounce, nil
}
