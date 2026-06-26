import { useQuery } from '@tanstack/react-query'
import * as dashboardApi from '@/api/dashboard'
import {
  mockDashboardIncidents,
  mockMonitoring,
  mockOverview,
  mockPerformance,
  mockRequestVolumeChart,
  mockResponseTimeChart,
  mockStatus,
  mockSuccessRate,
  mockUptime,
} from '@/mocks/data'

const USE_MOCK_FALLBACK = import.meta.env.VITE_USE_MOCK !== 'false'

async function withMockFallback<T>(fn: () => Promise<T>, mock: T): Promise<T> {
  try {
    return await fn()
  } catch {
    if (USE_MOCK_FALLBACK) return mock
    throw new Error('API unavailable')
  }
}

export function useDashboardOverview() {
  return useQuery({
    queryKey: ['dashboard', 'overview'],
    queryFn: () => withMockFallback(dashboardApi.overview, mockOverview),
  })
}

export function useDashboardStatus() {
  return useQuery({
    queryKey: ['dashboard', 'status'],
    queryFn: () => withMockFallback(dashboardApi.status, mockStatus),
  })
}

export function useDashboardMonitoring() {
  return useQuery({
    queryKey: ['dashboard', 'monitoring'],
    queryFn: () => withMockFallback(dashboardApi.monitoring, mockMonitoring),
  })
}

export function useDashboardPerformance() {
  return useQuery({
    queryKey: ['dashboard', 'performance'],
    queryFn: () => withMockFallback(dashboardApi.performance, mockPerformance),
  })
}

export function useDashboardSuccessRate() {
  return useQuery({
    queryKey: ['dashboard', 'success-rate'],
    queryFn: () => withMockFallback(dashboardApi.successRate, mockSuccessRate),
  })
}

export function useDashboardUptime() {
  return useQuery({
    queryKey: ['dashboard', 'uptime'],
    queryFn: () => withMockFallback(dashboardApi.uptime, mockUptime),
  })
}

export function useDashboardIncidents() {
  return useQuery({
    queryKey: ['dashboard', 'incidents'],
    queryFn: () => withMockFallback(dashboardApi.incidents, mockDashboardIncidents),
  })
}

// TODO(API): Replace mock chart data with GET /dashboard/history time-series when available
export function useResponseTimeChart() {
  return useQuery({
    queryKey: ['dashboard', 'response-time-chart'],
    queryFn: async () => mockResponseTimeChart,
  })
}

export function useRequestVolumeChart() {
  return useQuery({
    queryKey: ['dashboard', 'request-volume-chart'],
    queryFn: async () => mockRequestVolumeChart,
  })
}
