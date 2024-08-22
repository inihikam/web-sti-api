package main

import (
	"net/http"

	"github.com/inihikam/web-sti-api/config"
	"github.com/inihikam/web-sti-api/routes"
	"github.com/rs/cors"
)

func main() {
	config.LoadConfig()

	r := routes.SetupRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Authorization", "Content-Type"},
		ExposedHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 300,
	})

  	handler := c.Handler(r)

  	http.ListenAndServe(":8080", handler)
}
