import { apiClient } from './client'
import type {
  ApiResponse,
  DashboardIncident,
  Endpoint,
  EndpointCreateUpdate,
  EndpointMonitoring,
  EndpointStats,
} from '@/types/api'

export function normalizeEndpoint(raw: any): Endpoint {
  if (!raw) return raw
  return {
    id: raw.id ?? raw.ID ?? '',
    name: raw.name ?? raw.Name ?? '',
    url: raw.url ?? raw.URL ?? '',
    expected_status: raw.expected_status ?? raw.expectedStatus ?? raw.ExpectedStatus ?? 200,
    status: raw.status ?? undefined,
    last_checked: raw.last_checked ?? raw.lastCheckedAt ?? raw.LastCheckedAt ?? undefined,
    response_time: raw.response_time ?? undefined,
  }
}

export async function list(): Promise<Endpoint[]> {
  const { data } = await apiClient.get<ApiResponse<any[]>>('/endpoints')
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoints')
  return data.data.map(normalizeEndpoint)
}

export async function get(id: string | number): Promise<Endpoint> {
  const { data } = await apiClient.get<ApiResponse<any>>(`/endpoints/${id}`)
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoint')
  return normalizeEndpoint(data.data)
}

export async function create(payload: EndpointCreateUpdate): Promise<Endpoint> {
  const { data } = await apiClient.post<ApiResponse<any>>('/endpoints', payload)
  if (!data.success || !data.data) throw new Error('Failed to create endpoint')
  return normalizeEndpoint(data.data)
}

export async function update(id: string | number, payload: EndpointCreateUpdate): Promise<Endpoint> {
  const { data } = await apiClient.put<ApiResponse<any>>(`/endpoints/${id}`, payload)
  if (!data.success || !data.data) throw new Error('Failed to update endpoint')
  return normalizeEndpoint(data.data)
}


export async function remove(id: string | number): Promise<void> {
  const { data } = await apiClient.delete<ApiResponse<null>>(`/endpoints/${id}`)
  if (!data.success) throw new Error('Failed to delete endpoint')
}

export async function stats(id: string | number): Promise<EndpointStats> {
  const { data } = await apiClient.get<ApiResponse<any>>(`/endpoints/${id}/stats`)
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoint stats')
  return {
    average_response_time: data.data.average_latency ?? 0,
    min_response_time: 0,
    max_response_time: 0,
    success_rate: data.data.success_rate ?? 0,
    total_checks: data.data.total_checks ?? 0,
    uptime_percentage: data.data.success_rate ?? 0,
  }
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
