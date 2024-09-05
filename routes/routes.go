package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inihikam/web-sti-api/config"
	"github.com/inihikam/web-sti-api/controllers"
	"github.com/inihikam/web-sti-api/middlewares"
	"github.com/inihikam/web-sti-api/services"
)

func SetupRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Logger())
	r.Use(middlewares.CORS())

	taService := services.NewAnnouncementService(config.TaBaseURL)
	taController := controllers.NewAnnouncementController(taService)

	ta := r.Group("/api/ta")
	{
		ta.GET("/pengumuman", taController.GetAnnouncements)
		ta.GET("/pengumuman/:id", taController.GetAnnouncementDetail)
	}

	alumniService := services.NewAnnouncementService(config.AlumniBaseURL)
	alumniController := controllers.NewAnnouncementController(alumniService)

	alumni := r.Group("/api/alumni")
	{
		alumni.GET("/pengumuman", alumniController.GetAnnouncements)
		alumni.GET("/pengumuman/:id", alumniController.GetAnnouncementDetail)
		alumni.GET("/pengumumanLogang", alumniController.GetLogang)
		alumni.GET("/pengumumanLogang/:id", alumniController.GetLogangDetail)
	}

	bkService := services.NewAnnouncementService(config.BkBaseURL)
	bkController := controllers.NewAnnouncementController(bkService)

	bk := r.Group("/api/bk")
	{
		bk.GET("/pengumuman", bkController.GetAnnouncements)
		bk.GET("/pengumuman/:id", bkController.GetAnnouncementDetail)
	}

	kpService := services.NewAnnouncementService(config.KpBaseURL)
	kpController := controllers.NewAnnouncementController(kpService)

	kp := r.Group("/api/kp")
	{
		kp.GET("/pengumuman", kpController.GetAnnouncements)
		kp.GET("/pengumuman/:id", kpController.GetAnnouncementDetail)
	}

	r.GET("/api/dosen", controllers.GetLecturers)

	return r
}