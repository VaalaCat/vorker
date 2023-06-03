export interface LoginRequest {
	userName: string;
	password: string;
}

export interface LoginResponse {
	token: string;
	status: number;
}

export interface RegisterRequest {
	userName: string;
	email: string;
	password: string;
}

export interface RegisterResponse {
	status: number;
}

export interface UserInfo {
	userName: string;
	email: string;
	role: string;
}