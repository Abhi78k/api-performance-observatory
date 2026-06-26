import { Link } from 'react-router-dom'
import { AlertTriangle, CheckCircle2 } from 'lucide-react'
import {
  Badge,
  Card,
  EmptyState,
  ErrorState,
  Table,
  TableSkeleton,
  Typography,
} from '@/components/ui'
import { useIncidents } from '@/hooks/useIncidents'
import { formatDate, formatDuration, getStatusColor } from '@/utils/format'

export function IncidentsPage() {
  const { data, isLoading, isError, refetch } = useIncidents()

  const active = (data ?? []).filter((i) => !i.is_resolved)
  const historical = (data ?? []).filter((i) => i.is_resolved)

  const renderTable = (items: typeof data) => {
    if (!items?.length) {
      return (
        <EmptyState
          title="No incidents"
          description="All systems are operating normally."
        />
      )
    }

    return (
      <Table
        data={items}
        keyExtractor={(row) => row.id}
        columns={[
          {
            key: 'endpoint',
            header: 'Endpoint',
            render: (r) => (
              <Link to={`/endpoints/${r.endpoint_id}`} className="text-info hover:underline">
                {r.endpoint_name ?? `Endpoint #${r.endpoint_id}`}
              </Link>
            ),
          },
          {
            key: 'severity',
            header: 'Severity',
            render: (r) => (
              <Badge color={getStatusColor(r.severity ?? 'medium')}>
                {r.severity ?? 'medium'}
              </Badge>
            ),
          },
          {
            key: 'started_at',
            header: 'Started',
            render: (r) => formatDate(r.started_at),
          },
          {
            key: 'duration',
            header: 'Duration',
            render: (r) => {
              if (r.duration_minutes) return formatDuration(r.duration_minutes)
              if (r.resolved_at) {
                const mins = Math.floor(
                  (new Date(r.resolved_at).getTime() - new Date(r.started_at).getTime()) / 60000,
                )
                return formatDuration(mins)
              }
              const mins = Math.floor(
                (Date.now() - new Date(r.started_at).getTime()) / 60000,
              )
              return formatDuration(mins)
            },
          },
          {
            key: 'status',
            header: 'Status',
            render: (r) => (
              <Badge color={r.is_resolved ? 'success' : 'error'}>
                {r.is_resolved ? 'Resolved' : 'Active'}
              </Badge>
            ),
          },
        ]}
      />
    )
  }

  if (isLoading) return <TableSkeleton />
  if (isError) return <ErrorState onRetry={() => refetch()} />

  return (
    <div className="space-y-6">
      <div>
        <Typography variant="h5" color="white" fontWeight="bold">
          Incidents
        </Typography>
        <Typography variant="body2" color="text">
          Track active and historical incidents across your API infrastructure
        </Typography>
      </div>

      <Card>
        <div className="mb-4 flex items-center gap-2">
          <AlertTriangle className="h-5 w-5 text-error" />
          <Typography variant="lg" color="white" fontWeight="bold">
            Active Incidents
          </Typography>
          <Badge color="error">{active.length}</Badge>
        </div>
        {renderTable(active)}
      </Card>

      <Card>
        <div className="mb-4 flex items-center gap-2">
          <CheckCircle2 className="h-5 w-5 text-success" />
          <Typography variant="lg" color="white" fontWeight="bold">
            Historical Incidents
          </Typography>
          <Badge color="secondary">{historical.length}</Badge>
        </div>
        {renderTable(historical)}
      </Card>
    </div>
  )
}
