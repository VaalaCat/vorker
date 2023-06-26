package entities

type RegisterRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Status int `json:"status"`
}

func (r *RegisterRequest) Validate() bool {
	if r == nil {
		return false
	}
	if (r.UserName == "" && r.Email == "") || r.Password == "" {
		return false
	}
	if len(r.UserName) > 32 || len(r.Email) > 64 || len(r.Password) > 64 {
		return false
	}
	return true
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

func (l *LoginRequest) Validate() bool {
	if l == nil {
		return false
	}
	if l.UserName == "" || l.Password == "" {
		return false
	}
	if len(l.UserName) > 32 || len(l.Password) > 64 {
		return false
	}
	return true
}

type GetUserResponse struct {
	UserName string `json:"userName"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type DeleteWorkerRequest struct {
	UID string `json:"uid"`
}

func (d *DeleteWorkerRequest) Validate() bool {
	if d == nil {
		return false
	}
	if d.UID == "" {
		return false
	}
	if len(d.UID) > 64 {
		return false
	}
	return true
}

type AgentSyncWorkersReq struct {
	WorkerNames []string `json:"worker_names"`
}

type AgentSyncWorkersResp struct {
	WorkerList *WorkerList `json:"worker_list"`
}

type NotifyEventRequest struct {
	EventName string            `json:"event_name"`
	Extra     map[string][]byte `json:"extra"`
}

func (n *NotifyEventRequest) Validate() bool {
	if n == nil {
		return false
	}
	if n.EventName == "" {
		return false
	}
	if len(n.EventName) > 64 {
		return false
	}
	return true
}

type NotifyEventResponse struct {
	Status int `json:"status"` // 0: success, 1: failed
}

type SyncNodesResponse struct {
}

type RunWorkerResponse struct {
	Status  int    `json:"status"` // 0: success, 1: failed
	RunResp []byte `json:"run_resp"`
}
