package database

import (
	"time"
	"vorker/conf"
	"vorker/defs"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DBManager interface {
	GetDB(dbType string) *gorm.DB
	GetSqlite() *gorm.DB
	SetDB(dbType string, db *gorm.DB)
	WaitforInit()
	IsNil() bool
}

type DBManagerImpl struct {
	DBs map[string]*gorm.DB
	ch  chan bool
}

var (
	DBManagerInstance *DBManagerImpl
)

func InitDB() {
	DBManagerInstance = &DBManagerImpl{
		DBs: make(map[string]*gorm.DB),
		ch:  make(chan bool),
	}
	switch conf.AppConfigInstance.DBType {
	case defs.DBTypeSqlite:
		initSqlite()
	}
	close(DBManagerInstance.ch)
}

func GetManager() DBManager {
	return DBManagerInstance
}

func (m *DBManagerImpl) GetDB(dbType string) *gorm.DB {
	return m.DBs[dbType]
}

func (m *DBManagerImpl) GetSqlite() *gorm.DB {
	return m.DBs[defs.DBTypeSqlite]
}

func (m *DBManagerImpl) SetDB(dbType string, db *gorm.DB) {
	m.DBs[dbType] = db
}

func (m *DBManagerImpl) WaitforInit() {
	<-m.ch
}

func (m *DBManagerImpl) IsNil() bool {
	return m == nil
}

func GetDB() *gorm.DB {
	mgr := GetManager()
	for mgr.IsNil() {
		logrus.Infof("wait for init db")
		mgr = GetManager()
		time.Sleep(3 * time.Second)

	}
	mgr.WaitforInit()
	switch conf.AppConfigInstance.DBType {
	case defs.DBTypeSqlite:
		return mgr.GetSqlite()
	}
	return nil
}

// func CloseDB(db *gorm.DB) {
// 	tdb, err := db.DB()
// 	if err != nil {
// 		logrus.WithError(err).Errorf("Close DB error")
// 	}
// 	tdb.Close()
// }
