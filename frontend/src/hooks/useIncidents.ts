import { keepPreviousData, useQuery } from '@tanstack/react-query'
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

export function useIncidents(page?: number, limit?: number, isResolved?: string) {
  return useQuery({
    queryKey: ['incidents', page, limit, isResolved],
    queryFn: () =>
      withMockFallback(
        () => incidentsApi.list(page, limit, isResolved),
        {
          data: mockIncidents.filter((i) => {
            const matchesResolved = !isResolved || isResolved === 'all' || String(i.is_resolved) === isResolved
            return matchesResolved
          }).slice(((page ?? 1) - 1) * (limit ?? 10), (page ?? 1) * (limit ?? 10)),
          pagination: {
            page: page ?? 1,
            limit: limit ?? 10,
            totalItems: mockIncidents.length,
            totalPages: Math.ceil(mockIncidents.length / (limit ?? 10)),
            hasNext: (page ?? 1) * (limit ?? 10) < mockIncidents.length,
            hasPrevious: (page ?? 1) > 1,
          }
        }
      ),
    placeholderData: keepPreviousData,
  })
}

export function useActiveIncidents(page?: number, limit?: number) {
  return useQuery({
    queryKey: ['incidents', 'active', page, limit],
    queryFn: () =>
      withMockFallback(
        () => incidentsApi.active(page, limit),
        {
          data: mockIncidents.filter((i) => !i.is_resolved).slice(((page ?? 1) - 1) * (limit ?? 10), (page ?? 1) * (limit ?? 10)),
          pagination: {
            page: page ?? 1,
            limit: limit ?? 10,
            totalItems: mockIncidents.filter((i) => !i.is_resolved).length,
            totalPages: Math.ceil(mockIncidents.filter((i) => !i.is_resolved).length / (limit ?? 10)),
            hasNext: (page ?? 1) * (limit ?? 10) < mockIncidents.filter((i) => !i.is_resolved).length,
            hasPrevious: (page ?? 1) > 1,
          }
        }
      ),
    placeholderData: keepPreviousData,
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
