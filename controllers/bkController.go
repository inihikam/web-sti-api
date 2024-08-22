package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inihikam/web-sti-api/services"
)

func GetBkAnnouncements(c *gin.Context) {
	bkAnnounce, err := services.GetBkAnnounce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedAnnounce := make([]gin.H, len(bkAnnounce))

	for i, announce := range bkAnnounce {
		formattedAnnounce[i] = gin.H{
			"id": i+1,
			"title": announce.Title,
			"user": announce.User,
			"content": announce.Content,
			"date": announce.TanggalMulai,
		}
	}

	c.JSON(http.StatusOK, formattedAnnounce)
}