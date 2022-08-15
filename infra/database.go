package infra

import (
	"log"

	"github.com/guionardo/go-api-example/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase(config *Config) (db *gorm.DB, err error) {
	log.Printf("Connecting to database - %s", config.ConnectionString)
	return gorm.Open(sqlite.Open(config.ConnectionString), &gorm.Config{})
}

func ResetDatabase(db *gorm.DB) error {
	log.Printf("Resetting database")
	log.Printf("Removing feira table")
	res := db.Exec("DROP TABLE IF EXISTS feiras;")
	if res.Error != nil {
		return res.Error
	}
	log.Printf("Creating feira table")
	return RunMigration(db)
}

func RunMigration(db *gorm.DB) error {
	log.Printf("Running migrations")
	return db.AutoMigrate(&domain.Feira{})
}
