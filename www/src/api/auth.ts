import {
  UserInfo,
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
} from '@/types/body'
import api from './http'

export const login = async (req: LoginRequest) => {
  const res = await api.post('/api/auth/login', req)
  return res.data.data as LoginResponse
}

export const register = async (req: RegisterRequest) => {
  const res = await api.post('/api/auth/register', req)
  return res.data.data as RegisterResponse
}

export const getUserInfo = async () => {
  const res = await api.get('/api/user/info')
  return res.data.data as UserInfo
}

export const logout = () => {
  localStorage.removeItem('token')
  return api.get('/api/auth/logout')
}
