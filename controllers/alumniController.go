package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inihikam/web-sti-api/services"
)

func GetAlumniAnnouncements(c *gin.Context) {
	alumniAnnounce, err := services.GetAlumniAnnounce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedAnnounce := make([]gin.H, len(alumniAnnounce))

	for i, announce := range alumniAnnounce {
		formattedAnnounce[i] = gin.H{
			"id": i+1,
			"title": announce.Title,
			"user": announce.User,
			"content": announce.Content,
			"date": announce.Date.Format("2006-01-02 15:04:05"),
		}
	}

	c.JSON(http.StatusOK, formattedAnnounce)
}