package repository

import (
	"log"

	"github.com/guionardo/go-api-example/infra"
)

func RunReset() {
	log.Printf("Running Reset")
	config, err := infra.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %s", err)
	}
	service, err := infra.NewFeiraService(config)
	if err != nil {
		log.Fatalf("Error creating repository: %s", err)
	}
	err = service.Reset()
	if err != nil {
		log.Fatalf("Error reseting repository: %s", err)
	}
	feiras, err := ReadCsvFile("DEINFO_AB_FEIRASLIVRES_2014.csv")
	if err != nil {
		log.Fatalf("Error reading csv file: %s", err)
	}
	err = service.BulkSave(feiras)
	if err != nil {
		log.Fatalf("Error saving feira: %s", err)
	}

	log.Printf("Repository reset done - Saved %d feiras", len(feiras))
}
