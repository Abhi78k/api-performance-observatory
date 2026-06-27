import { apiClient } from './client'
import type { ApiResponse, HealthCheck } from '@/types/api'

export async function list(page?: number, limit?: number, endpointId?: string | number, success?: string): Promise<{ data: HealthCheck[]; pagination: any }> {
  let url = '/healthchecks'
  const params = new URLSearchParams()
  if (page) params.set('page', String(page))
  if (limit) params.set('limit', String(limit))
  if (endpointId && endpointId !== 'all') params.set('endpoint_id', String(endpointId))
  if (success && success !== 'all') {
    params.set('success', String(success === 'success'))
  }

  const queryStr = params.toString()
  if (queryStr) url += `?${queryStr}`

  const { data } = await apiClient.get<ApiResponse<HealthCheck[]>>(url)
  if (!data.success || !data.data) throw new Error('Failed to fetch health checks')
  return {
    data: data.data,
    pagination: (data as any).pagination,
  }
}

export async function byEndpoint(endpointId: string | number): Promise<HealthCheck[]> {
  const { data } = await apiClient.get<ApiResponse<HealthCheck[]>>(`/healthchecks/${endpointId}`)
  if (!data.success || !data.data) throw new Error('Failed to fetch health checks for endpoint')
  return data.data
}
