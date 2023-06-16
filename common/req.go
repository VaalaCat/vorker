package common

import "voker/defs"

type Request interface {
	*defs.DeleteWorkerRequest | *defs.LoginRequest | *defs.RegisterRequest
	Validate() bool
}
