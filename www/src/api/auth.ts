import { UserInfo, LoginRequest, LoginResponse, RegisterRequest, RegisterResponse } from "@/types/body"
import api from './http'

export const Login = (req: LoginRequest) => {
	return api.post('/api/auth/login', req).then((res) => res.data.data as LoginResponse)
}

export const Reg = (req: RegisterRequest) => {
	return api.post('/api/auth/register', req).then((res) => res.data.data as RegisterResponse)
}

export const GetUser = () => {
	return api.get('/api/user/info').then((res) => res.data.data as UserInfo)
}

export const Logout = () => {
	return api.get('/api/auth/logout')
}