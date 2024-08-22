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

	alumni := r.Group("/api/alumni")
	{
		alumni.GET("/pengumuman", controllers.GetAlumniAnnouncements)
	}

	bk := r.Group("/api/bk")
	{
		bk.GET("/pengumuman", controllers.GetBkAnnouncements)
	}

	return r
}