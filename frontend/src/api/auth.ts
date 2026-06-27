import { apiClient } from './client'
import type { ApiResponse, User } from '@/types/api'

export interface LoginPayload {
  email: string
  password: string
}

export interface RegisterPayload {
  email: string
  password: string
}

export async function login(payload: LoginPayload): Promise<ApiResponse<null>> {
  const { data } = await apiClient.post<ApiResponse<null>>('/auth/login', payload)
  if (!data.success) throw new Error(data.message ?? 'Login failed')
  return data
}

export async function register(payload: RegisterPayload): Promise<string> {
  const { data } = await apiClient.post<ApiResponse<null>>('/auth/register', payload)
  if (!data.success) throw new Error(data.message ?? 'Registration failed')
  return data.message ?? 'User registered successfully'
}

export async function me(): Promise<User> {
  const { data } = await apiClient.get<ApiResponse<User>>('/auth/me')
  if (!data.success || !data.data) throw new Error(data.message ?? 'Failed to fetch profile')
  return data.data
}

export async function logout(): Promise<void> {
  const { data } = await apiClient.post<ApiResponse<null>>('/auth/logout')
  if (!data.success) throw new Error(data.message ?? 'Logout failed')
}
