package api

import (
	"log"

	"github.com/guionardo/go-api-example/infra"
)

func RunAPI() {
	log.Printf("Running API")
	config, err := infra.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %s", err)
	}
	service, err := infra.NewFeiraService(config)
	if err != nil {
		log.Fatalf("Error creating repository: %s", err)
	}
	controller := &FeiraController{Service: service}
	server := NewServer(config)
	server.RegisterRoutes(controller)
	server.Start()
}
