import { useState } from "react";
import { useSearchParams } from "react-router-dom";
import {
  Badge,
  Card,
  EmptyState,
  ErrorState,
  Select,
  Table,
  TableSkeleton,
  Typography,
  Pagination,
} from "@/components/ui";
import { useHealthChecks } from "@/hooks/useHealthChecks";
import { useEndpoints } from "@/hooks/useEndpoints";
import { formatDate, formatMs } from "@/utils/format";

export function HealthChecksPage() {
  const [searchParams, setSearchParams] = useSearchParams();
  const page = Number(searchParams.get("page") ?? "1");

  const [successFilter, setSuccessFilter] = useState("all");
  const [endpointFilter, setEndpointFilter] = useState("all");

  const { data: endpointsResult } = useEndpoints(1, 100);
  const endpoints = endpointsResult?.data ?? [];

  const {
    data: checksResult,
    isLoading,
    isError,
    refetch,
  } = useHealthChecks(page, 10, endpointFilter, successFilter);
  const checks = checksResult?.data ?? [];
  const pagination = checksResult?.pagination;

  const handleSuccessFilterChange = (val: string) => {
    setSuccessFilter(val);
    setSearchParams((prev) => {
      prev.set("page", "1");
      return prev;
    });
  };

  const handleEndpointFilterChange = (val: string) => {
    setEndpointFilter(val);
    setSearchParams((prev) => {
      prev.set("page", "1");
      return prev;
    });
  };

  const handlePageChange = (newPage: number) => {
    setSearchParams((prev) => {
      prev.set("page", String(newPage));
      return prev;
    });
  };

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
              { value: "all", label: "All Results" },
              { value: "success", label: "Successful" },
              { value: "failed", label: "Failed" },
            ]}
            value={successFilter}
            onChange={(e) => handleSuccessFilterChange(e.target.value)}
            className="sm:w-48"
          />
          <Select
            label="Endpoint"
            options={[
              { value: "all", label: "All Endpoints" },
              ...endpoints.map((ep) => ({
                value: String(ep.id),
                label: ep.name,
              })),
            ]}
            value={endpointFilter}
            onChange={(e) => handleEndpointFilterChange(e.target.value)}
            className="sm:w-56"
          />
        </div>

        {isLoading && <TableSkeleton />}
        {isError && <ErrorState onRetry={() => refetch()} />}
        {!isLoading && !isError && checks.length === 0 && (
          <EmptyState
            title="No health checks found"
            description="Adjust your filters or wait for the next check cycle."
          />
        )}
        {!isLoading && !isError && checks.length > 0 && (
          <>
            <Table
              data={checks}
              keyExtractor={(row) => row.id}
              columns={[
                {
                  key: "endpoint",
                  header: "Endpoint",
                  render: (r) =>
                    r.endpoint_name && r.endpoint_name.trim() !== ""
                      ? r.endpoint_name
                      : `Endpoint #${r.endpoint_id}`,
                },
                {
                  key: "status_code",
                  header: "Status Code",
                  render: (r) => (
                    <Badge color={r.success ? "success" : "error"}>
                      {r.status_code == 0 ? 404 : r.status_code}
                    </Badge>
                  ),
                },
                {
                  key: "success",
                  header: "Success",
                  render: (r) => (
                    <Badge color={r.success ? "success" : "error"}>
                      {r.success ? "Yes" : "No"}
                    </Badge>
                  ),
                },
                {
                  key: "response_time",
                  header: "Response Time",
                  render: (r) => formatMs(r.response_time),
                },
                {
                  key: "checked_at",
                  header: "Timestamp",
                  render: (r) => formatDate(r.checked_at),
                },
              ]}
            />
            {pagination && (
              <Pagination
                page={page}
                totalPages={pagination.totalPages}
                onPageChange={handlePageChange}
                hasNext={pagination.hasNext}
                hasPrevious={pagination.hasPrevious}
              />
            )}
          </>
        )}
      </Card>
    </div>
  );
}
