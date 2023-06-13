package models

import (
	"voker/defs"
	"voker/utils/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	UID     string `gorm:"unique;not null"`
	Name    string `gorm:"unique;not null"`
	Secrect string `gorm:"not null"`
}

func init() {
	db := database.GetDB()
	defer database.CloseDB(db)
	db.AutoMigrate(&Node{})
	if err := db.FirstOrCreate(&Node{
		UID:  uuid.New().String(),
		Name: defs.DefaultNodeName,
	}).Error; err != nil {
		panic(err)
	}
}

func (Node) TableName() string {
	return "nodes"
}

func (n *Node) Create() error {
	db := database.GetDB()
	database.CloseDB(db)
	return db.Create(n).Error
}

func (n *Node) Update(uid string) error {
	db := database.GetDB()
	database.CloseDB(db)
	return db.Model(&Node{}).Where(
		&Node{
			UID: uid,
		},
	).Updates(
		&Node{
			UID:     n.UID,
			Name:    n.Name,
			Secrect: n.Secrect,
		},
	).Error
}

func (n *Node) Delete(uid string) error {
	db := database.GetDB()
	database.CloseDB(db)
	return db.Delete(&Node{
		UID: uid,
	}).Error
}

func AdminGetAllNodes() ([]*Node, error) {
	var nodes []*Node
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Find(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

func AdminGetAllNodesMap() (map[string]string, error) {
	nodes, err := AdminGetAllNodes()
	if err != nil {
		return nil, err
	}
	ans := make(map[string]string)
	for _, node := range nodes {
		ans[node.Name] = node.UID
	}
	return ans, nil
}
