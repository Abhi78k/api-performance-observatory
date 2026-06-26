import { useQuery } from '@tanstack/react-query'
import * as incidentsApi from '@/api/incidents'
import { mockIncidents } from '@/mocks/data'

const USE_MOCK_FALLBACK = import.meta.env.VITE_USE_MOCK !== 'false'

async function withMockFallback<T>(fn: () => Promise<T>, mock: T): Promise<T> {
  try {
    return await fn()
  } catch {
    if (USE_MOCK_FALLBACK) return mock
    throw new Error('API unavailable')
  }
}

export function useIncidents() {
  return useQuery({
    queryKey: ['incidents'],
    queryFn: () => withMockFallback(incidentsApi.list, mockIncidents),
  })
}

export function useActiveIncidents() {
  return useQuery({
    queryKey: ['incidents', 'active'],
    queryFn: () => withMockFallback(incidentsApi.active, mockIncidents.filter((i) => !i.is_resolved)),
  })
}

export function useIncident(id: string | number | undefined) {
  return useQuery({
    queryKey: ['incidents', id],
    queryFn: () =>
      withMockFallback(
        () => incidentsApi.get(id!),
        mockIncidents.find((i) => String(i.id) === String(id))!,
      ),
    enabled: !!id,
  })
}
