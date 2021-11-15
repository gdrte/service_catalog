//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"service_catalog/models"
)

func main() {
	// dsn := "host=localhost port=5432 user=postgres dbname=service_catalog password=postgres sslmode=disable "
	db, err := gorm.Open(sqlite.Open("service_catalog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	db.AutoMigrate(models.Service{}, models.Version{})
	seedServices(db, 40)
}

func seedVersions(howMany int) []models.Version {
	versions := make([]models.Version, howMany)
	for i := 0; i < howMany; i++ {
		versions[i] = models.Version{Ver: fmt.Sprintf("%d.%d.%d", rand.Intn(2), rand.Intn(10), rand.Intn(10))}
	}
	return versions
}
func seedServices(db *gorm.DB, howMany int) {
	for i := 0; i < howMany; i++ {
		svc := models.Service{
			Name:        gofakeit.Name(),
			Description: gofakeit.HackerPhrase(),
			Versions:    seedVersions(rand.Intn(4)),
		}
		db.Create(&svc)
	}
}
