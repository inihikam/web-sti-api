package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inihikam/web-sti-api/services"
)

type AnnouncementController struct {
	service *services.AnnouncementService
}

func NewAnnouncementController(service *services.AnnouncementService) *AnnouncementController {
    return &AnnouncementController{service: service}
}

func (ac *AnnouncementController) GetAnnouncements(c *gin.Context) {
    announcements, err := ac.service.GetAnnouncement()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }

    c.JSON(http.StatusOK, announcements)
}

func (ac *AnnouncementController) GetAnnouncementDetail(c *gin.Context) {
    id := c.Param("id")
    announcement, err := ac.service.GetAnnouncementDetail(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
	
    c.JSON(http.StatusOK, announcement)
}