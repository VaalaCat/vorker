package defs

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
