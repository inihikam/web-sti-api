package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inihikam/web-sti-api/models"
)

type AnnouncementService struct {
	baseURL string
}

func NewAnnouncementService(baseURL string) *AnnouncementService {
	return &AnnouncementService{baseURL: baseURL}
}

func (s *AnnouncementService) GetAnnouncement() ([]models.Announcement, error) {
	resp, err := http.Get(fmt.Sprintf("%s/pengumuman", s.baseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to get announcement: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	var announcements []models.Announcement
	err = json.NewDecoder(resp.Body).Decode(&announcements)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return announcements, nil
}

func (s *AnnouncementService) GetAnnouncementDetail(id string) (models.Announcement, error) {
    resp, err := http.Get(fmt.Sprintf("%s/pengumuman/%s", s.baseURL, id))
    if err != nil {
        return models.Announcement{}, fmt.Errorf("failed to fetch announcement detail: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return models.Announcement{}, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
    }

    var announcement models.Announcement
    err = json.NewDecoder(resp.Body).Decode(&announcement)
    if err != nil {
        return models.Announcement{}, fmt.Errorf("failed to decode announcement detail: %v", err)
    }

    return announcement, nil
}

func (s *AnnouncementService) GetLogang() ([]models.Logang, error) {
	resp, err := http.Get(fmt.Sprintf("%s/pengumumanLogang", s.baseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to get logang: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	var logang []models.Logang
	err = json.NewDecoder(resp.Body).Decode(&logang)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return logang, nil
}

func (s *AnnouncementService) GetLogangDetail(id string) (models.Logang, error) {
	resp, err := http.Get(fmt.Sprintf("%s/pengumumanLogang/%s", s.baseURL, id))
	if err != nil {
		return models.Logang{}, fmt.Errorf("failed to fetch logang detail: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Logang{}, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	var logang models.Logang
	err = json.NewDecoder(resp.Body).Decode(&logang)
	if err != nil {
		return models.Logang{}, fmt.Errorf("failed to decode logang detail: %v", err)
	}

	return logang, nil
}

