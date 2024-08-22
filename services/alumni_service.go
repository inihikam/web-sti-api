package services

import (
	"encoding/json"
	"net/http"

	"github.com/inihikam/web-sti-api/models"
)

func GetAlumniAnnounce() ([]models.AlumniAnnounce, error) {
	resp, err := http.Get("https://alumni-sti.inihikam.my.id/api/pengumuman")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var alumniAnnounce []models.AlumniAnnounce
	if err := json.NewDecoder(resp.Body).Decode(&alumniAnnounce); err != nil {
		return nil, err
	}

	return alumniAnnounce, nil
}
