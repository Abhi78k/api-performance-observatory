import { apiClient } from './client'
import type { ApiResponse, Incident } from '@/types/api'

export async function list(): Promise<Incident[]> {
  const { data } = await apiClient.get<ApiResponse<Incident[]>>('/incidents')
  if (!data.success || !data.data) throw new Error('Failed to fetch incidents')
  return data.data
}

export async function active(): Promise<Incident[]> {
  const { data } = await apiClient.get<ApiResponse<Incident[]>>('/incidents/active')
  if (!data.success || !data.data) throw new Error('Failed to fetch active incidents')
  return data.data
}

export async function get(id: string | number): Promise<Incident> {
  const { data } = await apiClient.get<ApiResponse<Incident>>(`/incidents/${id}`)
  if (!data.success || !data.data) throw new Error('Failed to fetch incident')
  return data.data
}
