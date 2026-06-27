export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
}

export interface User {
  id: string | number;
  email: string;
}

export interface LoginResponse {
  access_token: string;
}

export interface DashboardOverview {
  total_endpoints: number;
  monitored_endpoints: number;
  healthy_count: number;
  unhealthy_count: number;
}

export interface DashboardStatusItem {
  endpoint_id: string | number;
  endpoint_name: string;
  status: "healthy" | "unhealthy" | string;
  monitoring_duration_days: number;
}

export interface DashboardMonitoringItem {
  endpoint_id: string | number;
  endpoint_name: string;
  monitoring_started_at: string;
  monitoring_duration_days: number;
}

export interface DashboardPerformance {
  average_response_time: number;
  min_response_time: number;
  max_response_time: number;
}

export interface DashboardSuccessRate {
  total_checks: number;
  successful: number;
  failed: number;
  success_rate: number;
  failure_rate: number;
}

export interface DashboardUptime {
  uptime_percentage: number;
  total_incidents: number;
  total_downtime_minutes: number;
  average_incident_minutes: number;
}

export interface DashboardHistory {
  period: string;
  total_checks: number;
  average_response_time: number;
  success_rate: number;
}

export interface DashboardIncident {
  id: string | number;
  endpoint_id: string | number;
  started_at: string;
  resolved_at: string | null;
  is_resolved: boolean;
}

export interface Endpoint {
  id: string | number;
  name: string;
  url: string;
  expected_status: number;
  status?: "healthy" | "unhealthy" | "degraded" | string;
  last_checked?: string;
  response_time?: number;
}

export interface EndpointCreateUpdate {
  name: string;
  url: string;
  expected_status: number;
}

export interface EndpointStats {
  average_response_time: number;
  min_response_time: number;
  max_response_time: number;
  success_rate: number;
  total_checks: number;
  uptime_percentage: number;
}

export interface EndpointMonitoring {
  monitoring_started_at: string;
  monitoring_duration_days: number;
  check_interval_seconds: number;
}

export interface HealthCheck {
  id: string | number;
  endpoint_id: string | number;
  status_code: number;
  response_time: number;
  success: boolean;
  checked_at: string;
  endpoint_name?: string;
}

export interface Incident {
  id: string | number;
  endpoint_id: string | number;
  endpoint_name?: string;
  severity?: "critical" | "high" | "medium" | "low" | string;
  started_at: string;
  resolved_at: string | null;
  is_resolved: boolean;
  status?: "active" | "resolved" | string;
  duration_minutes?: number;
}

export type ArcType = "active" | "success" | "slow" | "failed";

export interface MonitoringNode {
  name: string;
  lat: number;
  lng: number;
}

export interface GlobeArc {
  start: MonitoringNode;
  end: MonitoringNode;
  type: ArcType;
}
