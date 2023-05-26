package database

import (
	"api-test/models"
	"api-test/models/seeds"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_HOST     string
	JWT_KEY     string
}

func InitDB(conf Config) *gorm.DB {
	conf = EnvDatabase()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("This is the error : ", err)
	}

	e := DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Assignment{}, &models.Answare{}, models.Discussions{}, &models.Question{})

	if e != nil {
		fmt.Println("This is the error : ", e)
	}

	seeds.NewSeeders(DB)
	return DB
}
