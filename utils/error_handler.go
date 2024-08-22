package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}