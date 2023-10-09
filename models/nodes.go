package models

import (
	"math/rand"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/utils/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	*entities.Node
}

func init() {
	db := database.GetDB()
	db.AutoMigrate(&Node{})
	if !conf.IsMaster() {
		return
	}
	if err := db.FirstOrCreate(&Node{
		Node: &entities.Node{
			UID:  uuid.New().String(),
			Name: defs.DefaultNodeName,
		},
	}).Error; err != nil {
		panic(err)
	}
	database.CloseDB(db)

	if self, err := GetNodeByNodeName(defs.DefaultNodeName); err != nil {
		panic(err)
	} else {
		conf.AppConfigInstance.NodeID = self.UID
	}
}

func (Node) TableName() string {
	return "nodes"
}

func (n *Node) Create() error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Create(n).Error
}

func (n *Node) Update(uid string) error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Model(&Node{}).Where(
		&Node{
			Node: &entities.Node{
				UID: uid,
			},
		},
	).Updates(
		&Node{
			Node: &entities.Node{
				UID:  n.UID,
				Name: n.Name,
			},
		},
	).Error
}

func (n *Node) Delete(uid string) error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Delete(&Node{
		Node: &entities.Node{
			UID: uid,
		},
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

// AdminGetAllNodesMap return a map with key is node name and value is node uid
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

func GetNodeByNodeName(nodeName string) (*Node, error) {
	node := Node{}
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(
		&Node{
			Node: &entities.Node{
				Name: nodeName}},
	).First(&node).Error; err != nil {
		return nil, err
	}
	return &node, nil
}

func GetAssignNode() (*Node, error) {
	nodes, err := AdminGetAllNodes()
	if err != nil {
		return nil, err
	}

	// random get a node
	idx := rand.Intn(len(nodes))
	return nodes[idx], nil
}

func NodeModels2Entities(nodes []*Node) []*entities.Node {
	ans := make([]*entities.Node, len(nodes))
	for i, node := range nodes {
		ans[i] = node.Node
	}
	return ans
}
