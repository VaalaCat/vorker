package models

import (
	"voker/entities"
	"voker/utils/database"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	*entities.Worker
}

func init() {
	db := database.GetDB()
	defer database.CloseDB(db)
	db.AutoMigrate(&Worker{})
}

func (w *Worker) TableName() string {
	return "workers"
}

func GetWorkerByUID(uid string) (*Worker, error) {
	var worker Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where("uid = ?", uid).First(&worker).Error; err != nil {
		return nil, err
	}
	return &worker, nil
}

func GetWorkersByNames(names []string) ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where("name in (?)", names).Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func GetAllWorkers() ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func Trans2Entities(workers []*Worker) []*entities.Worker {
	var entities []*entities.Worker
	for _, worker := range workers {
		entities = append(entities, worker.Worker)
	}
	return entities
}

func (w *Worker) Create() error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Create(w).Error
}

func (w *Worker) Update() error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Save(w).Error
}

func (w *Worker) Delete() error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Delete(w).Error
}
