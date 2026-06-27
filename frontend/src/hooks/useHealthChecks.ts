import { keepPreviousData, useQuery } from '@tanstack/react-query'
import * as healthchecksApi from '@/api/healthchecks'
import { mockHealthChecks } from '@/mocks/data'

const USE_MOCK_FALLBACK = import.meta.env.VITE_USE_MOCK !== 'false'

async function withMockFallback<T>(fn: () => Promise<T>, mock: T): Promise<T> {
  try {
    return await fn()
  } catch {
    if (USE_MOCK_FALLBACK) return mock
    throw new Error('API unavailable')
  }
}

export function useHealthChecks(page?: number, limit?: number, endpointId?: string | number, success?: string) {
  return useQuery({
    queryKey: ['healthchecks', page, limit, endpointId, success],
    queryFn: () =>
      withMockFallback(
        () => healthchecksApi.list(page, limit, endpointId, success),
        {
          data: mockHealthChecks.filter((h) => {
            const matchesSuccess = !success || success === 'all' || (success === 'success' && h.success) || (success === 'failed' && !h.success)
            const matchesEndpoint = !endpointId || endpointId === 'all' || String(h.endpoint_id) === String(endpointId)
            return matchesSuccess && matchesEndpoint
          }).slice(((page ?? 1) - 1) * (limit ?? 10), (page ?? 1) * (limit ?? 10)),
          pagination: {
            page: page ?? 1,
            limit: limit ?? 10,
            totalItems: mockHealthChecks.length,
            totalPages: Math.ceil(mockHealthChecks.length / (limit ?? 10)),
            hasNext: (page ?? 1) * (limit ?? 10) < mockHealthChecks.length,
            hasPrevious: (page ?? 1) > 1,
          }
        }
      ),
    placeholderData: keepPreviousData,
  })
}

export function useEndpointHealthChecks(endpointId: string | number | undefined) {
  return useQuery({
    queryKey: ['healthchecks', endpointId],
    queryFn: () =>
      withMockFallback(
        () => healthchecksApi.byEndpoint(endpointId!),
        mockHealthChecks.filter((h) => String(h.endpoint_id) === String(endpointId)),
      ),
    enabled: !!endpointId,
  })
}
