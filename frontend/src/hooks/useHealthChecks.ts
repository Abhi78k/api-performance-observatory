import { useQuery } from '@tanstack/react-query'
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

export function useHealthChecks() {
  return useQuery({
    queryKey: ['healthchecks'],
    queryFn: () => withMockFallback(healthchecksApi.list, mockHealthChecks),
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
