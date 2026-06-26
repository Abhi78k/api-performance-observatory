import { useParams, Link } from 'react-router-dom'
import { ArrowLeft, Clock, Globe, Shield } from 'lucide-react'
import {
  Badge,
  Button,
  Card,
  CardSkeleton,
  ErrorState,
  MiniStatisticsCard,
  Table,
  Typography,
} from '@/components/ui'
import { ChartCard } from '@/features/dashboard/ChartCard'
import { HealthCheckList, IncidentTimeline } from '@/features/dashboard/DashboardWidgets'
import {
  useEndpoint,
  useEndpointMonitoring,
  useEndpointStats,
} from '@/hooks/useEndpoints'
import { useEndpointHealthChecks } from '@/hooks/useHealthChecks'
import { mockIncidents, mockResponseTimeChart } from '@/mocks/data'
import { formatDate, formatMs, formatPercent, getStatusColor } from '@/utils/format'

export function EndpointDetailsPage() {
  const { id } = useParams<{ id: string }>()
  const endpoint = useEndpoint(id)
  const stats = useEndpointStats(id)
  const monitoring = useEndpointMonitoring(id)
  const healthChecks = useEndpointHealthChecks(id)

  const endpointIncidents = mockIncidents.filter(
    (i) => String(i.endpoint_id) === String(id),
  )

  if (endpoint.isLoading) return <CardSkeleton />
  if (endpoint.isError || !endpoint.data) {
    return <ErrorState onRetry={() => endpoint.refetch()} message="Failed to load endpoint details." />
  }

  const ep = endpoint.data

  return (
    <div className="space-y-6">
      <div className="flex items-start gap-4">
        <Link to="/endpoints">
          <Button variant="outlined" color="white" iconOnly size="small">
            <ArrowLeft className="h-4 w-4" />
          </Button>
        </Link>
        <div className="flex-1">
          <div className="flex flex-wrap items-center gap-3">
            <Typography variant="h5" color="white" fontWeight="bold">
              {ep.name}
            </Typography>
            <Badge color={getStatusColor(ep.status ?? 'unknown')}>{ep.status ?? 'unknown'}</Badge>
          </div>
          <Typography variant="body2" color="text" className="mt-1 break-all">
            {ep.url}
          </Typography>
        </div>
      </div>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
        <MiniStatisticsCard
          title="Avg Response Time"
          value={stats.data ? formatMs(stats.data.average_response_time) : '—'}
          icon={Clock}
        />
        <MiniStatisticsCard
          title="Success Rate"
          value={stats.data ? formatPercent(stats.data.success_rate) : '—'}
          icon={Shield}
          iconColor="#01B574"
        />
        <MiniStatisticsCard
          title="Uptime"
          value={stats.data ? formatPercent(stats.data.uptime_percentage) : '—'}
          icon={Globe}
        />
        <MiniStatisticsCard
          title="Total Checks"
          value={stats.data?.total_checks.toLocaleString() ?? '—'}
          icon={Clock}
        />
      </div>

      <div className="grid gap-4 lg:grid-cols-2">
        <Card>
          <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
            Overview
          </Typography>
          <dl className="space-y-3">
            <div className="flex justify-between">
              <Typography variant="caption" color="text">Expected Status</Typography>
              <Typography variant="button" color="white">{ep.expected_status}</Typography>
            </div>
            <div className="flex justify-between">
              <Typography variant="caption" color="text">Last Checked</Typography>
              <Typography variant="button" color="white">
                {ep.last_checked ? formatDate(ep.last_checked) : '—'}
              </Typography>
            </div>
            <div className="flex justify-between">
              <Typography variant="caption" color="text">Response Time</Typography>
              <Typography variant="button" color="white">
                {ep.response_time ? formatMs(ep.response_time) : '—'}
              </Typography>
            </div>
          </dl>
        </Card>

        <Card>
          <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
            Monitoring Information
          </Typography>
          {monitoring.isLoading ? (
            <CardSkeleton />
          ) : monitoring.data ? (
            <dl className="space-y-3">
              <div className="flex justify-between">
                <Typography variant="caption" color="text">Started At</Typography>
                <Typography variant="button" color="white">
                  {formatDate(monitoring.data.monitoring_started_at)}
                </Typography>
              </div>
              <div className="flex justify-between">
                <Typography variant="caption" color="text">Duration</Typography>
                <Typography variant="button" color="white">
                  {monitoring.data.monitoring_duration_days} days
                </Typography>
              </div>
              <div className="flex justify-between">
                <Typography variant="caption" color="text">Check Interval</Typography>
                <Typography variant="button" color="white">
                  {monitoring.data.check_interval_seconds}s
                </Typography>
              </div>
            </dl>
          ) : (
            <Typography variant="body2" color="text">No monitoring data available</Typography>
          )}
        </Card>
      </div>

      <ChartCard
        title="Response Time"
        subtitle="Performance over 24h"
        data={mockResponseTimeChart}
        dataKey="ms"
        unit="ms"
      />

      <div className="grid gap-4 lg:grid-cols-2">
        <HealthCheckList
          checks={healthChecks.data}
          isLoading={healthChecks.isLoading}
          isError={healthChecks.isError}
          onRetry={() => healthChecks.refetch()}
          limit={10}
        />
        <IncidentTimeline
          incidents={endpointIncidents}
          isLoading={false}
          isError={false}
          onRetry={() => {}}
          title="Incident History"
        />
      </div>

      {stats.data && (
        <Card>
          <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
            Performance Statistics
          </Typography>
          <Table
            data={[
              { metric: 'Min Response Time', value: formatMs(stats.data.min_response_time) },
              { metric: 'Max Response Time', value: formatMs(stats.data.max_response_time) },
              { metric: 'Average Response Time', value: formatMs(stats.data.average_response_time) },
              { metric: 'Success Rate', value: formatPercent(stats.data.success_rate) },
              { metric: 'Uptime', value: formatPercent(stats.data.uptime_percentage) },
            ]}
            keyExtractor={(r) => r.metric}
            columns={[
              { key: 'metric', header: 'Metric' },
              { key: 'value', header: 'Value', render: (r) => <span className="font-medium">{r.value}</span> },
            ]}
          />
        </Card>
      )}
    </div>
  )
}
