package database

import (
	"vorker/conf"
	"vorker/defs"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	godotenv.Load()
	switch conf.AppConfigInstance.DBType {
	case defs.DBTypeSqlite:
		initSqlite()
	}
}

func GetDB() *gorm.DB {
	switch conf.AppConfigInstance.DBType {
	case defs.DBTypeSqlite:
		return GetSqlite()
	}
	return nil
}

func CloseDB(db *gorm.DB) {
	tdb, err := db.DB()
	if err != nil {
		logrus.WithError(err).Errorf("Close DB error")
	}
	tdb.Close()
}
