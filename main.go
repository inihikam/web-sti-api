package main

import (
	"github.com/inihikam/web-sti-api/config"
	"github.com/inihikam/web-sti-api/routes"
)

func main() {
	config.LoadConfig()

	r := routes.SetupRouter()

	r.Run(":8080")
}
