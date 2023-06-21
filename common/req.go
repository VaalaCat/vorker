package common

import "vorker/defs"

type Request interface {
	*defs.DeleteWorkerRequest | *defs.LoginRequest | *defs.RegisterRequest
	Validate() bool
}
