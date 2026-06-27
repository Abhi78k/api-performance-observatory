import { keepPreviousData, useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import * as endpointsApi from '@/api/endpoints'
import * as dashboardApi from '@/api/dashboard'
import * as healthchecksApi from '@/api/healthchecks'
import { mockEndpoints } from '@/mocks/data'
import type { Endpoint, EndpointCreateUpdate } from '@/types/api'

const USE_MOCK_FALLBACK = import.meta.env.VITE_USE_MOCK !== 'false'

async function withMockFallback<T>(fn: () => Promise<T>, mock: T): Promise<T> {
  try {
    return await fn()
  } catch (err) {
    if (USE_MOCK_FALLBACK) return mock
    throw err
  }
}

export function useEndpoints(page?: number, limit?: number, search?: string, status?: string) {
  return useQuery({
    queryKey: ['endpoints', page, limit, search, status],
    queryFn: () =>
      withMockFallback(async () => {
        const result = await endpointsApi.list(page, limit, search, status)
        return result
      }, {
        data: mockEndpoints.filter((ep) => {
          const matchesSearch = !search || ep.name.toLowerCase().includes(search.toLowerCase()) || ep.url.toLowerCase().includes(search.toLowerCase())
          const matchesStatus = !status || status === 'all' || ep.status === status
          return matchesSearch && matchesStatus
        }).slice(((page ?? 1) - 1) * (limit ?? 10), (page ?? 1) * (limit ?? 10)),
        pagination: {
          page: page ?? 1,
          limit: limit ?? 10,
          totalItems: mockEndpoints.length,
          totalPages: Math.ceil(mockEndpoints.length / (limit ?? 10)),
          hasNext: (page ?? 1) * (limit ?? 10) < mockEndpoints.length,
          hasPrevious: (page ?? 1) > 1,
        }
      }),
    placeholderData: keepPreviousData,
  })
}

export function useEndpoint(id: string | number | undefined) {
  return useQuery({
    queryKey: ['endpoints', id],
    queryFn: () =>
      withMockFallback(async () => {
        const [ep, statusResult, healthChecks] = await Promise.all([
          endpointsApi.get(id!),
          dashboardApi.status(1, 100).catch(() => ({ data: [] })),
          healthchecksApi.byEndpoint(id!).catch(() => []),
        ])

        const statuses = 'data' in statusResult ? statusResult.data : statusResult
        const statusItem = Array.isArray(statuses) ? statuses.find((s: any) => String(s.endpoint_id) === String(ep.id)) : undefined
        const latestCheck = healthChecks.length > 0
          ? [...healthChecks].sort((a, b) => new Date(b.checked_at).getTime() - new Date(a.checked_at).getTime())[0]
          : undefined

        return {
          ...ep,
          status: statusItem?.status ?? (latestCheck?.success ? 'healthy' : latestCheck ? 'unhealthy' : 'unknown'),
          response_time: latestCheck?.response_time ?? undefined,
          last_checked: latestCheck?.checked_at ?? ep.last_checked,
        } as Endpoint
      }, mockEndpoints.find((e) => String(e.id) === String(id)) as Endpoint),
    enabled: !!id,
  })
}

export function useEndpointStats(id: string | number | undefined) {
  return useQuery({
    queryKey: ['endpoints', id, 'stats'],
    queryFn: () =>
      withMockFallback(async () => {
        const [stats, healthChecks] = await Promise.all([
          endpointsApi.stats(id!),
          healthchecksApi.byEndpoint(id!).catch(() => []),
        ])

        let min_response_time = 0
        let max_response_time = 0
        if (healthChecks.length > 0) {
          const latencies = healthChecks.map((c) => Number(c.response_time))
          min_response_time = Math.min(...latencies)
          max_response_time = Math.max(...latencies)
        }

        return {
          ...stats,
          min_response_time,
          max_response_time,
        }
      }, {
        average_response_time: 187,
        min_response_time: 42,
        max_response_time: 890,
        success_rate: 98.2,
        total_checks: 12400,
        uptime_percentage: 99.5,
      }),
    enabled: !!id,
  })
}

export function useEndpointMonitoring(id: string | number | undefined) {
  return useQuery({
    queryKey: ['endpoints', id, 'monitoring'],
    queryFn: () =>
      withMockFallback(async () => {
        const data = await endpointsApi.monitoring(id!)
        return {
          monitoring_started_at: data.monitoring_started_at,
          monitoring_duration_days: data.monitoring_duration_days,
          check_interval_seconds: data.check_interval_seconds ?? 60,
        }
      }, {
        monitoring_started_at: '2025-01-12T08:00:00Z',
        monitoring_duration_days: 45,
        check_interval_seconds: 60,
      }),
    enabled: !!id,
  })
}


export function useCreateEndpoint() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: (payload: EndpointCreateUpdate) => endpointsApi.create(payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['endpoints'] }),
  })
}

export function useUpdateEndpoint() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: ({ id, payload }: { id: string | number; payload: EndpointCreateUpdate }) =>
      endpointsApi.update(id, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['endpoints'] }),
  })
}

export function useDeleteEndpoint() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: (id: string | number) => endpointsApi.remove(id),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['endpoints'] }),
  })
}
