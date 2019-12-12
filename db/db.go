package db

import (
	"fmt"
	"gin-example-structure/config"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()
	dbInfo := fmt.Sprintf(" user=%s password=%s host=%s port=%s dbname=%s sslmode=%s ", c.Get("database.account"), c.Get("database.password"), c.Get("database.server"), c.Get("database.port"), c.Get("database.dbname"), c.Get("database.sslmode"))
	var err error
	db, err = gorm.Open("postgres", dbInfo)
	if err != nil {
		log.WithField("dbInfo", dbInfo).Fatal("connecting to database error")
	}
}

func GetDB() *gorm.DB {
	return db
}
