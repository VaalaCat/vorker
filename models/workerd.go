package models

import (
	"os"
	"path/filepath"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/utils"

	"vorker/utils/database"

	"github.com/google/uuid"
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

func GetWorkerByUID(userID uint, uid string) (*Worker, error) {
	var worker Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&Worker{
		Worker: &entities.Worker{
			UserID: uint64(userID),
		},
	}).Where(
		&Worker{
			Worker: &entities.Worker{
				UID: uid,
			},
		},
	).First(&worker).Error; err != nil {
		return nil, err
	}
	return &worker, nil
}

func GetWorkersByNames(userID uint, names []string) ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&Worker{
		Worker: &entities.Worker{
			UserID: uint64(userID),
		},
	}).Where("name in (?)", names).Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func AdminGetWorkersByNames(names []string) ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where("name in (?)", names).Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func GetAllWorkers(userID uint) ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&Worker{
		Worker: &entities.Worker{
			UserID: uint64(userID),
		},
	}).Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func AdminGetAllWorkers() ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func AdminGetWorkersByNodeName(nodeName string) ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&Worker{
		Worker: &entities.Worker{
			NodeName: nodeName,
		},
	}).Find(&workers).Error; err != nil {
		return nil, err
	}
	return workers, nil
}

func GetWorkers(userID uint, offset, limit int) ([]*Worker, error) {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&Worker{
		Worker: &entities.Worker{
			UserID: uint64(userID),
		},
	}).Offset(offset).Limit(limit).Find(&workers).Error; err != nil {
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

	err := w.UpdateFile()
	if err != nil {
		return err
	}

	return db.Create(w).Error
}

func (w *Worker) Update() error {
	db := database.GetDB()
	defer database.CloseDB(db)

	err := w.UpdateFile()
	if err != nil {
		return err
	}

	return db.Save(w).Error
}

func (w *Worker) Delete() error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Delete(w).Error
}

func (w *Worker) Flush() error {
	port, err := utils.GetAvailablePort(defs.DefaultHostName)
	if err != nil {
		return err
	}
	if len(w.TunnelID) == 0 {
		w.TunnelID = uuid.New().String()
	}

	if err = w.DeleteFile(); err != nil {
		return err
	}

	w.Port = int32(port)
	if err = w.Update(); err != nil {
		return err
	}

	if err = w.UpdateFile(); err != nil {
		return err
	}
	return nil
}

func (w *Worker) DeleteFile() error {
	return os.RemoveAll(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.WorkerCodePath,
			w.UID,
		),
	)
}

func (w *Worker) UpdateFile() error {
	return utils.WriteFile(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.WorkerCodePath,
			w.UID,
			w.Entry),
		string(w.Code))
}

func SyncWorkers(workerList *entities.WorkerList) error {
	var workers []*Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	// write all workers to db

	UIDs := []string{}
	for _, worker := range workerList.Workers {
		UIDs = append(UIDs, worker.UID)
		workers = append(workers, &Worker{
			Worker: worker,
		})
	}

	if err := db.Create(workers).Error; err != nil {
		return err
	}

	// delete workers not in workerList
	if err := db.Where("uid not in (?)", UIDs).Delete(&Worker{}).Error; err != nil {
		return err
	}

	return nil
}

func SyncAddWorker(worker *entities.Worker) error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Create(&Worker{
		Worker: worker,
	}).Error
}

func SyncDeleteWorker(worker *entities.Worker) error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Where(&Worker{
		Worker: worker,
	}).Delete(&Worker{}).Error
}

func SyncUpdateWorker(workerName string, worker *entities.Worker) error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Model(&Worker{}).Where(&Worker{
		Worker: &entities.Worker{
			Name: workerName,
		},
	}).Updates(&Worker{
		Worker: worker,
	}).Error
}
