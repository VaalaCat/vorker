package model

import (
	"log"
	"voker/conf"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	var err error
	godotenv.Load()
	if conf.AppConfigInstance.DBType != "sqlite" {
		return
	}

	dbPath := conf.AppConfigInstance.DBPath
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Panic(err, "Initializing DB Error")
	}
	CloseDB(db)
}

func GetSqlite() *gorm.DB {
	dbPath := conf.AppConfigInstance.DBPath
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

func CloseDB(db *gorm.DB) {
	tdb, err := db.DB()
	if err != nil {
		logrus.WithError(err).Errorf("Close DB error")
	}
	tdb.Close()
}
