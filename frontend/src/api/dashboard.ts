import { apiClient } from './client'
import type {
  ApiResponse,
  DashboardHistory,
  DashboardIncident,
  DashboardMonitoringItem,
  DashboardOverview,
  DashboardPerformance,
  DashboardStatusItem,
  DashboardSuccessRate,
  DashboardUptime,
} from '@/types/api'

export async function overview(): Promise<DashboardOverview> {
  const { data } = await apiClient.get<ApiResponse<DashboardOverview>>('/dashboard/overview')
  if (!data.success || !data.data) throw new Error('Failed to fetch dashboard overview')
  return data.data
}

export async function status(page?: number, limit?: number): Promise<{ data: DashboardStatusItem[]; pagination: any }> {
  let url = '/dashboard/status'
  const params = new URLSearchParams()
  if (page) params.set('page', String(page))
  if (limit) params.set('limit', String(limit))

  const queryStr = params.toString()
  if (queryStr) url += `?${queryStr}`

  const { data } = await apiClient.get<ApiResponse<DashboardStatusItem[]>>(url)
  if (!data.success || !data.data) throw new Error('Failed to fetch endpoint status')
  return {
    data: data.data,
    pagination: (data as any).pagination,
  }
}

export async function monitoring(): Promise<DashboardMonitoringItem[]> {
  const { data } = await apiClient.get<ApiResponse<DashboardMonitoringItem[]>>('/dashboard/monitoring')
  if (!data.success || !data.data) throw new Error('Failed to fetch monitoring data')
  return data.data
}

export async function performance(): Promise<DashboardPerformance> {
  const { data } = await apiClient.get<ApiResponse<DashboardPerformance>>('/dashboard/performance')
  if (!data.success || !data.data) throw new Error('Failed to fetch performance data')
  return data.data
}

export async function successRate(): Promise<DashboardSuccessRate> {
  const { data } = await apiClient.get<ApiResponse<DashboardSuccessRate>>('/dashboard/success-rate')
  if (!data.success || !data.data) throw new Error('Failed to fetch success rate')
  return data.data
}

export async function uptime(): Promise<DashboardUptime> {
  const { data } = await apiClient.get<ApiResponse<DashboardUptime>>('/dashboard/uptime')
  if (!data.success || !data.data) throw new Error('Failed to fetch uptime data')
  return data.data
}

export async function history(): Promise<DashboardHistory> {
  const { data } = await apiClient.get<ApiResponse<DashboardHistory>>('/dashboard/history')
  if (!data.success || !data.data) throw new Error('Failed to fetch history data')
  return data.data
}

export async function incidents(): Promise<DashboardIncident[]> {
  const { data } = await apiClient.get<ApiResponse<DashboardIncident[]>>('/dashboard/incidents')
  if (!data.success || !data.data) throw new Error('Failed to fetch dashboard incidents')
  return data.data
}
