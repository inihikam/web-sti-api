package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inihikam/web-sti-api/services"
)

func GetKpAnnouncements(c *gin.Context) {
	kpAnnounce, err := services.GetKpAnnounce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kpAnnounce)
}