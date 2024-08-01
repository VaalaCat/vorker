package models

import (
	"context"
	"vorker/utils/database"

	"gorm.io/gorm"
)

type WorkerVersion struct {
	gorm.Model
	UID      string
	WorkerID string
	Name     string
	FileID   string
}

func (w *WorkerVersion) TableName() string {
	return "worker_versions"
}

func GetFileByVersionUID(c context.Context, versionID string) (*File, error) {
	var version WorkerVersion
	var file File
	db := database.GetDB()
	if err := db.Where(&WorkerVersion{UID: versionID}).First(&version).Error; err != nil {
		return nil, err
	}

	if err := db.Where(&File{UID: version.FileID}).First(&file).Error; err != nil {
		return nil, err
	}

	return &file, nil
}
