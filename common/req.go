package common

import (
	"vorker/entities"
)

type Request interface {
	*entities.DeleteWorkerRequest | *entities.LoginRequest | *entities.RegisterRequest |
		*entities.NotifyEventRequest
	Validate() bool
}
