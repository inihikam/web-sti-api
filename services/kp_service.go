package services

import (
	"encoding/json"
	"net/http"

	"github.com/inihikam/web-sti-api/models"
)

func GetKpAnnounce() ([]models.KpAnnounce, error) {
	resp, err := http.Get("https://kp-sti.inihikam.my.id/api/pengumuman")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var kpAnnounce []models.KpAnnounce
	if err := json.NewDecoder(resp.Body).Decode(&kpAnnounce); err != nil {
		return nil, err
	}

	return kpAnnounce, nil
}