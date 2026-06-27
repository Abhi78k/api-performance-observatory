import { apiClient } from './client'
import type { ApiResponse, Incident } from '@/types/api'

export async function list(page?: number, limit?: number, isResolved?: string): Promise<{ data: Incident[]; pagination: any }> {
  let url = '/incidents'
  const params = new URLSearchParams()
  if (page) params.set('page', String(page))
  if (limit) params.set('limit', String(limit))
  if (isResolved !== undefined && isResolved !== 'all') params.set('is_resolved', isResolved)

  const queryStr = params.toString()
  if (queryStr) url += `?${queryStr}`

  const { data } = await apiClient.get<ApiResponse<Incident[]>>(url)
  if (!data.success || !data.data) throw new Error('Failed to fetch incidents')
  return {
    data: data.data,
    pagination: (data as any).pagination,
  }
}

export async function active(page?: number, limit?: number): Promise<{ data: Incident[]; pagination: any }> {
  let url = '/incidents/active'
  const params = new URLSearchParams()
  if (page) params.set('page', String(page))
  if (limit) params.set('limit', String(limit))

  const queryStr = params.toString()
  if (queryStr) url += `?${queryStr}`

  const { data } = await apiClient.get<ApiResponse<Incident[]>>(url)
  if (!data.success || !data.data) throw new Error('Failed to fetch active incidents')
  return {
    data: data.data,
    pagination: (data as any).pagination,
  }
}

export async function get(id: string | number): Promise<Incident> {
  const { data } = await apiClient.get<ApiResponse<Incident>>(`/incidents/${id}`)
  if (!data.success || !data.data) throw new Error('Failed to fetch incident')
  return data.data
}
