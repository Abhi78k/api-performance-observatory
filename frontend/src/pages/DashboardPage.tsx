import { useState } from "react";
import {
  Activity,
  AlertTriangle,
  CheckCircle2,
  Clock,
  Globe,
  Server,
  TrendingUp,
  XCircle,
  GlobeX,
  GlobeOff,
  ClockAlert,
} from "lucide-react";
import { MonitoringGlobe } from "@/components/Globe/MonitoringGlobe";
import {
  MiniStatisticsCard,
  Card,
  Typography,
  CardSkeleton,
  ErrorState,
} from "@/components/ui";
import { ChartCard } from "@/features/dashboard/ChartCard";
import {
  EndpointStatusList,
  HealthCheckList,
  IncidentTimeline,
  RankedEndpoints,
} from "@/features/dashboard/DashboardWidgets";
import {
  useDashboardIncidents,
  useDashboardOverview,
  useDashboardPerformance,
  useDashboardStatus,
  useDashboardSuccessRate,
  useDashboardUptime,
  useRequestVolumeChart,
  useResponseTimeChart,
} from "@/hooks/useDashboard";
import { useHealthChecks } from "@/hooks/useHealthChecks";
import { useActiveIncidents } from "@/hooks/useIncidents";
import { useEndpoints } from "@/hooks/useEndpoints";
import { formatMs, formatPercent } from "@/utils/format";

export function DashboardPage() {
  const [statusPage, setStatusPage] = useState(1);

  const overview = useDashboardOverview();
  const status = useDashboardStatus(statusPage, 10);
  const performance = useDashboardPerformance();
  const successRate = useDashboardSuccessRate();
  const uptime = useDashboardUptime();
  const incidents = useDashboardIncidents();
  const activeIncidents = useActiveIncidents();
  const healthChecks = useHealthChecks();
  const responseChart = useResponseTimeChart();
  const volumeChart = useRequestVolumeChart();
  const endpointsQuery = useEndpoints(1, 100);

  const endpoints = endpointsQuery.data?.data ?? [];

  const slowest = [...endpoints]
    .sort((a, b) => (b.response_time ?? 0) - (a.response_time ?? 0))
    .slice(0, 5)
    .map((e) => ({
      id: e.id,
      name: e.name,
      value: formatMs(e.response_time ?? 0),
    }));

  const errorRates = endpoints
    .filter((e) => e.status === "unhealthy" || e.status === "degraded")
    .slice(0, 5)
    .map((e) => ({
      id: e.id,
      name: e.name,
      value: e.status === "unhealthy" ? "4.2%" : "1.8%",
    }));

  if (overview.isLoading) {
    return (
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {Array.from({ length: 4 }).map((_, i) => (
          <CardSkeleton key={i} />
        ))}
      </div>
    );
  }

  if (overview.isError) {
    return <ErrorState onRetry={() => overview.refetch()} />;
  }

  const data = overview.data!;

  return (
    <div className="space-y-4">
      {/*<div>
        <Typography variant="h5" color="white" fontWeight="bold">
          Global Overview
        </Typography>
        <Typography variant="body2" color="text">
          Real-time API monitoring across {data.monitored_endpoints} endpoints
          worldwide
        </Typography>
      </div>*/}

      <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <MiniStatisticsCard
          title="Total Endpoints"
          value={data.total_endpoints}
          icon={Server}
        />
        <MiniStatisticsCard
          title="Healthy"
          value={data.healthy_count}
          subtitle={`${formatPercent((data.healthy_count / data.total_endpoints) * 100)}`}
          subtitleColor="success"
          icon={CheckCircle2}
          iconColor="#01B574"
        />
        <MiniStatisticsCard
          title="Unhealthy"
          value={data.unhealthy_count}
          subtitleColor="error"
          icon={XCircle}
          iconColor="#E31A1A"
        />
        <MiniStatisticsCard
          title="Overall Uptime"
          value={
            uptime.data ? formatPercent(uptime.data.uptime_percentage) : "—"
          }
          subtitle={
            uptime.data ? `${uptime.data.total_incidents} incidents` : undefined
          }
          subtitleColor="info"
          icon={TrendingUp}
        />
      </div>

      <div className="grid gap-4 lg:grid-cols-12">
        <div className="lg:col-span-3 space-y-4">
          <MiniStatisticsCard
            title="Avg Response Time"
            value={
              performance.data
                ? formatMs(performance.data.average_response_time)
                : "—"
            }
            subtitle={
              performance.data
                ? `min ${formatMs(performance.data.min_response_time)}`
                : undefined
            }
            icon={Clock}
          />
          <MiniStatisticsCard
            title="Request Volume"
            value={
              successRate.data
                ? successRate.data.total_checks.toLocaleString()
                : "—"
            }
            subtitle={
              successRate.data
                ? formatPercent(successRate.data.success_rate) + " success"
                : undefined
            }
            icon={Activity}
          />
          <MiniStatisticsCard
            title="Active Incidents"
            value={activeIncidents.data?.data?.length ?? 0}
            subtitleColor="error"
            icon={AlertTriangle}
            iconColor="#E31A1A"
          />
          <MiniStatisticsCard
            title="Total Endpoints"
            value={data.total_endpoints}
            icon={Server}
          />
          <MiniStatisticsCard
            title="Failure Rate"
            value={
              successRate.data
                ? formatPercent(successRate.data.failure_rate)
                : "—"
            }
            icon={GlobeOff}
            iconColor="#01B574"
          />
        </div>

        <div className="lg:col-span-6">
          <Card className="relative h-[420px] overflow-hidden !p-0 mb-4.5">
            <div className="absolute inset-x-0 top-4 z-10 text-center">
              <Typography
                variant="subtitle1"
                color="white"
                fontWeight="bold"
                className="flex items-center justify-center gap-2"
              >
                <Globe className="h-4 w-4 text-info" />
                Global Monitoring Network
              </Typography>
              <Typography variant="caption" color="text">
                Live traffic across 8 monitoring regions
              </Typography>
            </div>
            <MonitoringGlobe className="h-full w-full" />
          </Card>
          <div className="grid gap-4 lg:col-span-3">
            <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-2">
              <MiniStatisticsCard
                title="Downtime"
                value={uptime.data ? uptime.data.total_downtime_minutes + "m" : "—"}
                subtitleColor="error"
                icon={GlobeX}
                iconColor="#E31A1A"
              />
              <MiniStatisticsCard
                title="Avg Incident Duration"
                value={uptime.data ? uptime.data.average_incident_minutes.toFixed(1) + "m" : "—"}
                subtitle={
                  uptime.data
                    ? `${uptime.data.total_incidents} incidents`
                    : undefined
                }
                subtitleColor="info"
                icon={ClockAlert}
              />
            </div>
          </div>
        </div>

        <div className="lg:col-span-3 space-y-4">
          <RankedEndpoints
            title="Top Slowest Endpoints"
            items={slowest.splice(0, 4)}
          />
          <RankedEndpoints
            title="Top Error Rates"
            items={
              errorRates.splice(0, 1).length
                ? errorRates
                : slowest.slice(0, 3).map((e) => ({ ...e, value: "0.5%" }))
            }
          />
        </div>
      </div>

      <div className="grid gap-4 md:grid-cols-2">
        <ChartCard
          title="Response Time"
          subtitle="Average latency over 24h"
          data={responseChart.data ?? []}
          dataKey="ms"
          unit="ms"
        />
        <ChartCard
          title="Request Volume"
          subtitle="Total requests over 24h"
          data={volumeChart.data ?? []}
          dataKey="requests"
          type="bar"
          color="#01B574"
        />
      </div>

      <div className="grid gap-4 lg:grid-cols-3">
        <EndpointStatusList
          items={status.data?.data}
          isLoading={status.isLoading}
          isError={status.isError}
          onRetry={() => status.refetch()}
          page={statusPage}
          totalPages={status.data?.pagination?.totalPages}
          onPageChange={setStatusPage}
          hasNext={status.data?.pagination?.hasNext}
          hasPrevious={status.data?.pagination?.hasPrevious}
        />
        <HealthCheckList
          checks={healthChecks.data?.data}
          isLoading={healthChecks.isLoading}
          isError={healthChecks.isError}
          onRetry={() => healthChecks.refetch()}
        />
        <IncidentTimeline
          incidents={incidents.data}
          isLoading={incidents.isLoading}
          isError={incidents.isError}
          onRetry={() => incidents.refetch()}
        />
      </div>

      {/*{uptime.data && (
        <Card>
          <Typography
            variant="lg"
            color="white"
            fontWeight="bold"
            className="mb-4"
          >
            KPI Summary
          </Typography>
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
            <div>
              <Typography variant="caption" color="text">
                Uptime
              </Typography>
              <Typography variant="h6" color="success" fontWeight="bold">
                {formatPercent(uptime.data.uptime_percentage)}
              </Typography>
            </div>
            <div>
              <Typography variant="caption" color="text">
                Total Downtime
              </Typography>
              <Typography variant="h6" color="white" fontWeight="bold">
                {uptime.data.total_downtime_minutes.toFixed(2)}m
              </Typography>
            </div>
            <div>
              <Typography variant="caption" color="text">
                Avg Incident Duration
              </Typography>
              <Typography variant="h6" color="white" fontWeight="bold">
                {uptime.data.average_incident_minutes.toFixed(1)}m
              </Typography>
            </div>
            <div>
              <Typography variant="caption" color="text">
                Failure Rate
              </Typography>
              <Typography variant="h6" color="error" fontWeight="bold">
                {successRate.data
                  ? formatPercent(successRate.data.failure_rate)
                  : "—"}
              </Typography>
            </div>
          </div>
        </Card>
      )}*/}
    </div>
  );
}
