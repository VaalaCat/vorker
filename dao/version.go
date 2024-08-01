package dao

import (
	"context"
	"fmt"
	"vorker/models"
	"vorker/utils/database"
)

func NewWorkerVersion(version *models.WorkerVersion) error {
	if version == nil {
		return fmt.Errorf("version is nil")
	}

	db := database.GetDB()
	return db.Save(version).Error
}

func GetVersionByUID(c context.Context, versionID string) (*models.WorkerVersion, error) {
	var version models.WorkerVersion
	db := database.GetDB()
	if err := db.Where(&models.WorkerVersion{UID: versionID}).First(&version).Error; err != nil {
		return nil, err
	}
	return &version, nil
}
