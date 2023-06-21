package database

import (
	"vorker/conf"
	"vorker/defs"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func initSqlite() {
	var err error
	godotenv.Load()
	if conf.AppConfigInstance.DBType != defs.DBTypeSqlite {
		return
	}

	dbPath := conf.AppConfigInstance.DBPath
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logrus.Panic(err, "Initializing DB Error")
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
