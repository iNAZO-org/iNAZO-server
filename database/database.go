package database

import (
	"fmt"
	"karintou8710/iNAZO-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(models interface{}) {
	c := config.GetConfig()
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		c.GetString("db.host"),
		c.GetString("db.user"),
		c.GetString("db.password"),
		c.GetString("db.dbname"),
		c.GetString("db.port"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(models)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db

}
