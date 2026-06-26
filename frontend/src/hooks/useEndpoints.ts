import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import * as endpointsApi from '@/api/endpoints'
import { mockEndpoints } from '@/mocks/data'
import type { EndpointCreateUpdate } from '@/types/api'

const USE_MOCK_FALLBACK = import.meta.env.VITE_USE_MOCK !== 'false'

async function withMockFallback<T>(fn: () => Promise<T>, mock: T): Promise<T> {
  try {
    return await fn()
  } catch {
    if (USE_MOCK_FALLBACK) return mock
    throw new Error('API unavailable')
  }
}

export function useEndpoints() {
  return useQuery({
    queryKey: ['endpoints'],
    queryFn: () => withMockFallback(endpointsApi.list, mockEndpoints),
  })
}

export function useEndpoint(id: string | number | undefined) {
  return useQuery({
    queryKey: ['endpoints', id],
    queryFn: () =>
      withMockFallback(
        () => endpointsApi.get(id!),
        mockEndpoints.find((e) => String(e.id) === String(id))!,
      ),
    enabled: !!id,
  })
}

export function useEndpointStats(id: string | number | undefined) {
  return useQuery({
    queryKey: ['endpoints', id, 'stats'],
    queryFn: () =>
      withMockFallback(
        () => endpointsApi.stats(id!),
        {
          average_response_time: 187,
          min_response_time: 42,
          max_response_time: 890,
          success_rate: 98.2,
          total_checks: 12400,
          uptime_percentage: 99.5,
        },
      ),
    enabled: !!id,
  })
}

export function useEndpointMonitoring(id: string | number | undefined) {
  return useQuery({
    queryKey: ['endpoints', id, 'monitoring'],
    queryFn: () =>
      withMockFallback(
        () => endpointsApi.monitoring(id!),
        {
          monitoring_started_at: '2025-01-12T08:00:00Z',
          monitoring_duration_days: 45,
          check_interval_seconds: 60,
        },
      ),
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
