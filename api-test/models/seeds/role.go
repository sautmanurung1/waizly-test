package seeds

import (
	"api-test/models"

	"gorm.io/gorm"
)

func roleSeeder(db *gorm.DB) Seed {
	seeds := []models.Role{
		{
			ID:   1,
			Name: "teacher",
		},
		{
			ID:   2,
			Name: "student",
		},
	}

	model := &models.Role{}

	return Seed{
		entity: model,
		run: func(db *gorm.DB) (err error) {
			for _, v := range seeds {
				err = db.Create(&v).Error
			}
			return err
		},
	}
}
