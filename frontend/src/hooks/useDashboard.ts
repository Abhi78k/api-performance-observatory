import { useQuery } from '@tanstack/react-query'
import * as dashboardApi from '@/api/dashboard'
import * as healthchecksApi from '@/api/healthchecks'
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

export function useResponseTimeChart() {
  return useQuery({
    queryKey: ['dashboard', 'response-time-chart'],
    queryFn: async () => {
      try {
        const checks = await healthchecksApi.list()
        if (!checks || checks.length === 0) return mockResponseTimeChart

        const now = new Date()
        const intervals = Array.from({ length: 7 }).map((_, i) => {
          const time = new Date(now.getTime() - (6 - i) * 4 * 60 * 60 * 1000)
          const label = time.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
          return { time: label, msSum: 0, count: 0, date: time }
        })

        for (const check of checks) {
          const checkTime = new Date(check.checked_at).getTime()
          let closestIdx = 0
          let minDiff = Infinity
          intervals.forEach((interval, idx) => {
            const diff = Math.abs(interval.date.getTime() - checkTime)
            if (diff < minDiff) {
              minDiff = diff
              closestIdx = idx
            }
          })

          if (minDiff < 4 * 60 * 60 * 1000) {
            intervals[closestIdx].msSum += check.response_time
            intervals[closestIdx].count++
          }
        }

        return intervals.map((interval, i) => ({
          time: i === 6 ? 'Now' : interval.time,
          ms: interval.count > 0 ? Math.round(interval.msSum / interval.count) : 100, // fallback average
        }))
      } catch {
        return mockResponseTimeChart
      }
    },
  })
}

export function useRequestVolumeChart() {
  return useQuery({
    queryKey: ['dashboard', 'request-volume-chart'],
    queryFn: async () => {
      try {
        const checks = await healthchecksApi.list()
        if (!checks || checks.length === 0) return mockRequestVolumeChart

        const now = new Date()
        const intervals = Array.from({ length: 7 }).map((_, i) => {
          const time = new Date(now.getTime() - (6 - i) * 4 * 60 * 60 * 1000)
          const label = time.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
          return { time: label, count: 0, date: time }
        })

        for (const check of checks) {
          const checkTime = new Date(check.checked_at).getTime()
          let closestIdx = 0
          let minDiff = Infinity
          intervals.forEach((interval, idx) => {
            const diff = Math.abs(interval.date.getTime() - checkTime)
            if (diff < minDiff) {
              minDiff = diff
              closestIdx = idx
            }
          })

          if (minDiff < 4 * 60 * 60 * 1000) {
            intervals[closestIdx].count++
          }
        }

        return intervals.map((interval, i) => ({
          time: i === 6 ? 'Now' : interval.time,
          requests: interval.count,
        }))
      } catch {
        return mockRequestVolumeChart
      }
    },
  })
}

