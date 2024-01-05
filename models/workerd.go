package models

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/rpc"
	"vorker/tunnel"
	"vorker/utils"
	"vorker/utils/database"

	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	*entities.Worker
}

func init() {
	go func() {
		if conf.AppConfigInstance.LitefsEnabled {
			if !conf.IsMaster() {
				return
			}
			utils.WaitForPort("localhost", conf.AppConfigInstance.LitefsPrimaryPort)
		}
		db := database.GetDB()
		defer database.CloseDB(db)
		for err := db.AutoMigrate(&Worker{}); err != nil; err = db.AutoMigrate(&Worker{}) {
			logrus.WithError(err).Errorf("auto migrate worker error, sleep 5s and retry")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		if conf.AppConfigInstance.LitefsEnabled {
			utils.WaitForPort("localhost", conf.AppConfigInstance.LitefsPrimaryPort)
		}
		NodeWorkersInit()
	}()
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

func AdminGetWorkerByName(name string) (*Worker, error) {
	var worker Worker
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(
		&Worker{
			Worker: &entities.Worker{
				Name: name,
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

func AdminGetAllWorkersTunnelMap() (map[string]string, error) {
	workers, err := AdminGetAllWorkers()
	if err != nil {
		return nil, err
	}
	tunnelMap := make(map[string]string)
	for _, worker := range workers {
		tunnelMap[worker.Name] = worker.TunnelID
	}
	return tunnelMap, nil
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
	if w.NodeName == conf.AppConfigInstance.NodeName {
		port, err := utils.GetAvailablePort(defs.DefaultHostName)
		if err != nil {
			return err
		}
		w.Port = int32(port)
		tunnel.GetClient().Add(w.GetUID(), utils.WorkerHostPrefix(w.GetName()), int(w.GetPort()))
		err = w.UpdateFile()
		if err != nil {
			return err
		}
		if !conf.IsMaster() && conf.AppConfigInstance.LitefsEnabled {
			return nil
		}
	} else {
		n, err := GetNodeByNodeName(w.NodeName)
		if err != nil {
			return err
		}
		wp, err := proto.Marshal(w)
		if err != nil {
			return err
		}
		go rpc.EventNotify(n.Node, defs.EventAddWorker, map[string][]byte{
			defs.KeyWorkerProto: wp,
		})
	}

	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Create(w).Error
}

func (w *Worker) Update() error {
	if w.NodeName == conf.AppConfigInstance.NodeName {
		port, err := utils.GetAvailablePort(defs.DefaultHostName)
		if err != nil {
			return err
		}
		w.Port = int32(port)
		tunnel.GetClient().Delete(w.GetUID())
		tunnel.GetClient().Add(w.GetUID(),
			utils.WorkerHostPrefix(w.GetName()), port)
		err = w.UpdateFile()
		if err != nil {
			return err
		}
	}
	if !conf.IsMaster() && conf.AppConfigInstance.LitefsEnabled {
		return nil
	}
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Where(&Worker{
		Worker: &entities.Worker{
			UID: w.UID,
		},
	}).Save(w).Error
}

func (w *Worker) Delete() error {
	if w.NodeName == conf.AppConfigInstance.NodeName {
		tunnel.GetClient().Delete(w.GetUID())
	} else {
		n, err := GetNodeByNodeName(w.NodeName)
		if err != nil {
			return err
		}
		wp, err := proto.Marshal(w)
		if err != nil {
			return err
		}
		go rpc.EventNotify(n.Node, defs.EventDeleteWorker, map[string][]byte{
			defs.KeyWorkerProto: wp,
		})
	}
	if err := w.DeleteFile(); err != nil {
		return err
	}

	if !conf.IsMaster() && conf.AppConfigInstance.LitefsEnabled {
		return nil
	}
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Unscoped().Where(
		&Worker{Worker: &entities.Worker{
			UID: w.UID,
		}}).Delete(&Worker{}).Error
}

func (w *Worker) Flush() error {
	if w.NodeName != conf.AppConfigInstance.NodeName {
		n, err := GetNodeByNodeName(w.NodeName)
		if err != nil {
			return err
		}
		wp, err := proto.Marshal(w)
		if err != nil {
			return err
		}
		return rpc.EventNotify(n.Node, defs.EventFlushWorker, map[string][]byte{
			defs.KeyWorkerProto: wp,
		})
	}
	if len(w.TunnelID) == 0 {
		w.TunnelID = uuid.New().String()
	}

	if err := w.DeleteFile(); err != nil {
		return err
	}

	if err := w.Update(); err != nil {
		return err
	}
	return nil
}

func (w *Worker) DeleteFile() error {
	return os.RemoveAll(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.WorkerInfoPath,
			w.UID,
		),
	)
}

func (w *Worker) UpdateFile() error {
	return utils.WriteFile(
		filepath.Join(
			conf.AppConfigInstance.WorkerdDir,
			defs.WorkerInfoPath,
			w.UID,
			defs.WorkerCodePath,
			w.Entry),
		string(w.Code))
}

func (w *Worker) Run() ([]byte, error) {
	var addr string
	if w.GetNodeName() == conf.AppConfigInstance.NodeName {
		addr = fmt.Sprintf("http://%s:%d", w.GetHostName(), w.GetPort())
	} else {
		addr = fmt.Sprintf("http://%s:%d", conf.AppConfigInstance.TunnelHost,
			conf.AppConfigInstance.TunnelEntryPort)
	}
	resp, err := req.C().R().SetHeader(
		defs.HeaderHost, fmt.Sprintf("%s%s", w.Name, conf.AppConfigInstance.WorkerURLSuffix),
	).Get(addr)
	if err != nil || resp == nil || resp.StatusCode >= 299 {
		return nil, err
	}
	return resp.Bytes(), nil
}

func SyncWorkers(workerList *entities.WorkerList) error {
	partialFail := false
	UIDs := []string{}
	oldWorkers, err := AdminGetWorkersByNodeName(conf.AppConfigInstance.NodeName)
	if err != nil {
		return err
	}

	oldWorkerUIDMap := lo.SliceToMap(oldWorkers, func(w *Worker) (string, bool) { return w.UID, true })

	for _, worker := range workerList.Workers {
		modelWorker := &Worker{Worker: worker}
		UIDs = append(UIDs, worker.UID)
		if _, ok := oldWorkerUIDMap[worker.UID]; ok {
			continue
		}

		if err := modelWorker.Delete(); err != nil && err != gorm.ErrRecordNotFound {
			logrus.WithError(err).Errorf("sync workers db delete error, worker is: %+v", worker)
			partialFail = true
			continue
		}

		if err := modelWorker.Create(); err != nil {
			logrus.WithError(err).Errorf("sync workers db create error, worker is: %+v", worker)
			partialFail = true
			continue
		}

		if err := modelWorker.DeleteFile(); err != nil {
			logrus.WithError(err).Errorf("sync workers delete file error, worker is: %+v", worker)
			partialFail = true
			continue
		}

		if err := modelWorker.UpdateFile(); err != nil {
			logrus.WithError(err).Errorf("sync workers update file error, worker is: %+v", worker)
			partialFail = true
			continue
		}
		if err := utils.GenWorkerConfig(modelWorker.Worker); err != nil {
			logrus.WithError(err).Errorf("sync workers gen config error, worker is: %+v", worker)
			partialFail = true
			continue
		}
	}

	// delete workers that not in workerList
	for _, worker := range oldWorkers {
		if !utils.ContainsString(UIDs, worker.UID) {
			if err := worker.Delete(); err != nil {
				logrus.WithError(err).Errorf("sync workers delete worker error, worker is: %+v", worker)
				partialFail = true
				continue
			}
			if err := worker.DeleteFile(); err != nil {
				logrus.WithError(err).Errorf("sync workers delete file error, worker is: %+v", worker)
				partialFail = true
				continue
			}
		}
	}

	if partialFail {
		return errors.New("partial fail")
	}

	return nil
}
