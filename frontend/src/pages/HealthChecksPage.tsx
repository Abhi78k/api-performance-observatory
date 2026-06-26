import { useState } from 'react'
import {
  Badge,
  Card,
  EmptyState,
  ErrorState,
  Select,
  Table,
  TableSkeleton,
  Typography,
} from '@/components/ui'
import { useHealthChecks } from '@/hooks/useHealthChecks'
import { formatDate, formatMs } from '@/utils/format'

export function HealthChecksPage() {
  const { data, isLoading, isError, refetch } = useHealthChecks()
  const [successFilter, setSuccessFilter] = useState('all')
  const [endpointFilter, setEndpointFilter] = useState('all')

  const endpoints = [...new Set((data ?? []).map((h) => h.endpoint_name ?? String(h.endpoint_id)))]

  const filtered = (data ?? []).filter((check) => {
    const matchesSuccess =
      successFilter === 'all' ||
      (successFilter === 'success' && check.success) ||
      (successFilter === 'failed' && !check.success)
    const matchesEndpoint =
      endpointFilter === 'all' ||
      (check.endpoint_name ?? String(check.endpoint_id)) === endpointFilter
    return matchesSuccess && matchesEndpoint
  })

  return (
    <div className="space-y-6">
      <div>
        <Typography variant="h5" color="white" fontWeight="bold">
          Health Checks
        </Typography>
        <Typography variant="body2" color="text">
          View all health check results across your endpoints
        </Typography>
      </div>

      <Card>
        <div className="mb-4 flex flex-col gap-3 sm:flex-row">
          <Select
            label="Status"
            options={[
              { value: 'all', label: 'All Results' },
              { value: 'success', label: 'Successful' },
              { value: 'failed', label: 'Failed' },
            ]}
            value={successFilter}
            onChange={(e) => setSuccessFilter(e.target.value)}
            className="sm:w-48"
          />
          <Select
            label="Endpoint"
            options={[
              { value: 'all', label: 'All Endpoints' },
              ...endpoints.map((ep) => ({ value: ep, label: ep })),
            ]}
            value={endpointFilter}
            onChange={(e) => setEndpointFilter(e.target.value)}
            className="sm:w-56"
          />
        </div>

        {isLoading && <TableSkeleton />}
        {isError && <ErrorState onRetry={() => refetch()} />}
        {!isLoading && !isError && filtered.length === 0 && (
          <EmptyState title="No health checks found" description="Adjust your filters or wait for the next check cycle." />
        )}
        {!isLoading && !isError && filtered.length > 0 && (
          <Table
            data={filtered}
            keyExtractor={(row) => row.id}
            columns={[
              {
                key: 'endpoint',
                header: 'Endpoint',
                render: (r) => r.endpoint_name ?? `Endpoint #${r.endpoint_id}`,
              },
              {
                key: 'status_code',
                header: 'Status Code',
                render: (r) => (
                  <Badge color={r.success ? 'success' : 'error'}>{r.status_code}</Badge>
                ),
              },
              {
                key: 'success',
                header: 'Success',
                render: (r) => (
                  <Badge color={r.success ? 'success' : 'error'}>
                    {r.success ? 'Yes' : 'No'}
                  </Badge>
                ),
              },
              {
                key: 'response_time',
                header: 'Response Time',
                render: (r) => formatMs(r.response_time),
              },
              {
                key: 'checked_at',
                header: 'Timestamp',
                render: (r) => formatDate(r.checked_at),
              },
            ]}
          />
        )}
      </Card>
    </div>
  )
}
