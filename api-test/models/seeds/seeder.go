package seeds

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Seed struct {
	entity interface{}
	run    func(*gorm.DB) error
}

func NewSeeders(db *gorm.DB) {
	for _, v := range seedAll(db) {
		if db.Migrator().HasTable(&v.entity) {
			err := db.First(&v.entity).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := v.run(db); err != nil {
					log.Fatalf("Running Seed Error %e", err)
				}
			}
		}
	}
}

func seedAll(db *gorm.DB) []Seed {
	return []Seed{
		// Add New Seeder in Here
		roleSeeder(db),
	}
}
