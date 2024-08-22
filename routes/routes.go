package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inihikam/web-sti-api/controllers"
	"github.com/inihikam/web-sti-api/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Logger())

	ta := r.Group("/api/ta")
	{
		ta.GET("/pengumuman", controllers.GetTaAnnouncements)
	}

	return r
}