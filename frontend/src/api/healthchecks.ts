import { apiClient } from './client'
import type { ApiResponse, HealthCheck } from '@/types/api'

export async function list(): Promise<HealthCheck[]> {
  const { data } = await apiClient.get<ApiResponse<HealthCheck[]>>('/healthchecks')
  if (!data.success || !data.data) throw new Error('Failed to fetch health checks')
  return data.data
}

export async function byEndpoint(endpointId: string | number): Promise<HealthCheck[]> {
  const { data } = await apiClient.get<ApiResponse<HealthCheck[]>>(`/healthchecks/${endpointId}`)
  if (!data.success || !data.data) throw new Error('Failed to fetch health checks for endpoint')
  return data.data
}
