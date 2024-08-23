package main

import (
	"log"

	"github.com/inihikam/web-sti-api/config"
	"github.com/inihikam/web-sti-api/routes"
)

func main() {
	config := config.LoadConfig()

	r := routes.SetupRouter(config)

	log.Println("Server started at :8080")
	r.Run(":8080")
}
