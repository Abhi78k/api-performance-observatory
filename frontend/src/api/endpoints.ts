import { apiClient } from './client'
import type {
  ApiResponse,
  DashboardIncident,
  Endpoint,
  EndpointCreateUpdate,
  EndpointMonitoring,
  EndpointStats,
} from '@/types/api'

export async function list(): Promise<Endpoint[]> {
  const { data } = await apiClient.get<ApiResponse<Endpoint[]>>('/endpoints')
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoints')
  return data.data
}

export async function get(id: string | number): Promise<Endpoint> {
  const { data } = await apiClient.get<ApiResponse<Endpoint>>(`/endpoints/${id}`)
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoint')
  return data.data
}

export async function create(payload: EndpointCreateUpdate): Promise<Endpoint> {
  const { data } = await apiClient.post<ApiResponse<Endpoint>>('/endpoints', payload)
  if (!data.success || !data.data) throw new Error('Failed to create endpoint')
  return data.data
}

export async function update(id: string | number, payload: EndpointCreateUpdate): Promise<Endpoint> {
  const { data } = await apiClient.put<ApiResponse<Endpoint>>(`/endpoints/${id}`, payload)
  if (!data.success || !data.data) throw new Error('Failed to update endpoint')
  return data.data
}

export async function remove(id: string | number): Promise<void> {
  const { data } = await apiClient.delete<ApiResponse<null>>(`/endpoints/${id}`)
  if (!data.success) throw new Error('Failed to delete endpoint')
}

export async function stats(id: string | number): Promise<EndpointStats> {
  const { data } = await apiClient.get<ApiResponse<EndpointStats>>(`/endpoints/${id}/stats`)
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoint stats')
  return data.data
}

export async function monitoring(id: string | number): Promise<EndpointMonitoring> {
  const { data } = await apiClient.get<ApiResponse<EndpointMonitoring>>(`/endpoints/${id}/monitoring`)
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoint monitoring')
  return data.data
}

export async function incident(id: string | number): Promise<DashboardIncident | null> {
  const { data } = await apiClient.get<ApiResponse<DashboardIncident | null>>(`/endpoints/${id}/incidents`)
  if (!data.success) throw new Error('Failed to fetch endpoint incident')
  return data.data ?? null
}
