package model

import (
	"voker/entities"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	entities.Worker
}

