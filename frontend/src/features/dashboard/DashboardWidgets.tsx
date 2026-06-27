import { Link } from "react-router-dom";
import {
  Card,
  Badge,
  Typography,
  ErrorState,
  CardSkeleton,
  Button,
} from "@/components/ui";
import { formatDate, getStatusColor } from "@/utils/format";
import type {
  DashboardIncident,
  DashboardStatusItem,
  HealthCheck,
  Incident,
} from "@/types/api";

interface EndpointStatusListProps {
  items?: DashboardStatusItem[];
  isLoading: boolean;
  isError: boolean;
  onRetry: () => void;
  page?: number;
  totalPages?: number;
  onPageChange?: (page: number) => void;
  hasNext?: boolean;
  hasPrevious?: boolean;
}

export function EndpointStatusList({
  items,
  isLoading,
  isError,
  onRetry,
  page,
  totalPages,
  onPageChange,
  hasNext,
  hasPrevious,
}: EndpointStatusListProps) {
  if (isLoading) return <CardSkeleton />;
  if (isError)
    return (
      <Card>
        <ErrorState onRetry={onRetry} />
      </Card>
    );

  return (
    <Card>
      <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
        Endpoint Status
      </Typography>
      <div className="space-y-3">
        {items?.map((item) => (
          <Link
            key={item.endpoint_id}
            to={`/endpoints/${item.endpoint_id}`}
            className="flex items-center justify-between rounded-lg bg-white/5 px-4 py-2 transition-colors hover:bg-white/10"
          >
            <div>
              <Typography variant="button" color="white" className="pr-2">
                {item.endpoint_name}
              </Typography>
              <Typography variant="caption" color="text">
                {item.monitoring_duration_days.toFixed(2)} days monitored
              </Typography>
            </div>
            <Badge color={getStatusColor(item.status)}>{item.status}</Badge>
          </Link>
        ))}
      </div>
      {totalPages && totalPages > 1 && onPageChange && (
        <div className="mt-4 flex items-center justify-between border-t border-white/10 pt-3">
          <Typography variant="caption" color="text">
            Page {page} of {totalPages}
          </Typography>
          <div className="flex gap-2">
            <Button
              variant="outlined"
              size="small"
              disabled={!hasPrevious}
              onClick={() => onPageChange(page! - 1)}
              className="h-8 px-2 py-0 text-xs"
            >
              Prev
            </Button>
            <Button
              variant="outlined"
              size="small"
              disabled={!hasNext}
              onClick={() => onPageChange(page! + 1)}
              className="h-8 px-2 py-0 text-xs"
            >
              Next
            </Button>
          </div>
        </div>
      )}
    </Card>
  );
}

interface IncidentTimelineProps {
  incidents?: (DashboardIncident | Incident)[];
  endpointNames?: Record<string | number, string>;
  isLoading: boolean;
  isError: boolean;
  onRetry: () => void;
  title?: string;
}

export function IncidentTimeline({
  incidents,
  endpointNames = {},
  isLoading,
  isError,
  onRetry,
  title = "Recent Incidents",
}: IncidentTimelineProps) {
  if (isLoading) return <CardSkeleton />;
  if (isError)
    return (
      <Card>
        <ErrorState onRetry={onRetry} />
      </Card>
    );

  return (
    <Card>
      <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
        {title}
      </Typography>
      {!incidents?.length ? (
        <Typography variant="body2" color="text">
          No incidents recorded
        </Typography>
      ) : (
        <div className="space-y-3">
          {incidents.map((inc, idx) => {
            const name =
              "endpoint_name" in inc && inc.endpoint_name && inc.endpoint_name.trim() !== ""
                ? inc.endpoint_name
                : (endpointNames[inc.endpoint_id] ??
                  `Endpoint #${inc.endpoint_id}`);
            return (
              <div
                key={`${inc.id}-${idx}`}
                className="flex items-start gap-3 border-l-2 border-error/50 pl-3"
              >
                <div className="flex-1">
                  <Typography variant="button" color="white">
                    {name}
                  </Typography>
                  <Typography variant="caption" color="text" className="block">
                    {formatDate(inc.started_at)}
                    {inc.is_resolved && inc.resolved_at
                      ? ` → ${formatDate(inc.resolved_at)}`
                      : " → Active"}
                  </Typography>
                </div>
                <Badge color={inc.is_resolved ? "success" : "error"}>
                  {inc.is_resolved ? "Resolved" : "Active"}
                </Badge>
              </div>
            );
          })}
        </div>
      )}
    </Card>
  );
}

interface HealthCheckListProps {
  checks?: HealthCheck[];
  isLoading: boolean;
  isError: boolean;
  onRetry: () => void;
  limit?: number;
}

export function HealthCheckList({
  checks,
  isLoading,
  isError,
  onRetry,
  limit = 5,
}: HealthCheckListProps) {
  if (isLoading) return <CardSkeleton />;
  if (isError)
    return (
      <Card>
        <ErrorState onRetry={onRetry} />
      </Card>
    );

  const display = checks?.slice(0, limit) ?? [];

  return (
    <Card>
      <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
        Recent Health Checks
      </Typography>
      <div className="space-y-2">
        {display.map((check) => (
          <div
            key={check.id}
            className="flex items-center justify-between rounded-lg bg-white/5 px-3 py-2"
          >
            <div>
              <Typography variant="button" color="white" className="pr-2">
                {check.endpoint_name && check.endpoint_name.trim() !== "" ? check.endpoint_name : `Endpoint #${check.endpoint_id}`}
              </Typography>
              <Typography variant="caption" color="text">
                {formatDate(check.checked_at)}
              </Typography>
            </div>
            <div className="flex items-center gap-2">
              <Typography variant="caption" color="text">
                {check.response_time}ms
              </Typography>
              <Badge color={check.success ? "success" : "error"}>
                {check.status_code}
              </Badge>
            </div>
          </div>
        ))}
      </div>
    </Card>
  );
}

interface RankedEndpointsProps {
  title: string;
  items: { name: string; value: string; id: string | number }[];
}

export function RankedEndpoints({ title, items }: RankedEndpointsProps) {
  return (
    <Card>
      <Typography variant="lg" color="white" fontWeight="bold" className="mb-4">
        {title}
      </Typography>
      <div className="space-y-3 flex flex-col gap-[1.5px]">
        {items.map((item, i) => (
          <Link
            key={item.id}
            to={`/endpoints/${item.id}`}
            className="flex items-center gap-3 rounded-lg bg-white/5 px-3 py-2 transition-colors hover:bg-white/10"
          >
            <span className="flex h-6 w-6 items-center justify-center rounded-full bg-info/20 text-xs font-bold text-info">
              {i + 1}
            </span>
            <Typography
              variant="button"
              color="white"
              className="flex-1 truncate"
            >
              {item.name}
            </Typography>
            <Typography variant="caption" color="text">
              {item.value}
            </Typography>
          </Link>
        ))}
      </div>
    </Card>
  );
}
