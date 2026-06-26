import type {
  DashboardIncident,
  DashboardMonitoringItem,
  DashboardOverview,
  DashboardPerformance,
  DashboardStatusItem,
  DashboardSuccessRate,
  DashboardUptime,
  Endpoint,
  GlobeArc,
  HealthCheck,
  Incident,
  MonitoringNode,
} from '@/types/api'

export const MONITORING_NODES: MonitoringNode[] = [
  { name: 'New York', lat: 40.7128, lng: -74.006 },
  { name: 'London', lat: 51.5074, lng: -0.1278 },
  { name: 'Frankfurt', lat: 50.1109, lng: 8.6821 },
  { name: 'Mumbai', lat: 19.076, lng: 72.8777 },
  { name: 'Singapore', lat: 1.3521, lng: 103.8198 },
  { name: 'Tokyo', lat: 35.6762, lng: 139.6503 },
  { name: 'Sydney', lat: -33.8688, lng: 151.2093 },
  { name: 'São Paulo', lat: -23.5505, lng: -46.6333 },
]

// TODO(API): Replace mock globe arcs with backend location/traffic metadata when available
export const MOCK_GLOBE_ARCS: GlobeArc[] = [
  { start: MONITORING_NODES[0], end: MONITORING_NODES[1], type: 'success' },
  { start: MONITORING_NODES[2], end: MONITORING_NODES[4], type: 'active' },
  { start: MONITORING_NODES[3], end: MONITORING_NODES[5], type: 'slow' },
  { start: MONITORING_NODES[6], end: MONITORING_NODES[7], type: 'failed' },
  { start: MONITORING_NODES[1], end: MONITORING_NODES[3], type: 'success' },
  { start: MONITORING_NODES[4], end: MONITORING_NODES[0], type: 'active' },
]

export const mockOverview: DashboardOverview = {
  total_endpoints: 24,
  monitored_endpoints: 22,
  healthy_count: 19,
  unhealthy_count: 3,
}

export const mockPerformance: DashboardPerformance = {
  average_response_time: 187,
  min_response_time: 42,
  max_response_time: 2340,
}

export const mockSuccessRate: DashboardSuccessRate = {
  total_checks: 145820,
  successful: 143612,
  failed: 2208,
  success_rate: 98.48,
  failure_rate: 1.52,
}

export const mockUptime: DashboardUptime = {
  uptime_percentage: 99.87,
  total_incidents: 12,
  total_downtime_minutes: 94,
  average_incident_minutes: 7.8,
}

export const mockStatus: DashboardStatusItem[] = [
  { endpoint_id: 1, endpoint_name: 'Auth API', status: 'healthy', monitoring_duration_days: 45 },
  { endpoint_id: 2, endpoint_name: 'Payments Gateway', status: 'unhealthy', monitoring_duration_days: 30 },
  { endpoint_id: 3, endpoint_name: 'User Service', status: 'healthy', monitoring_duration_days: 60 },
  { endpoint_id: 4, endpoint_name: 'Search Index', status: 'degraded', monitoring_duration_days: 22 },
  { endpoint_id: 5, endpoint_name: 'Notifications', status: 'healthy', monitoring_duration_days: 15 },
]

export const mockMonitoring: DashboardMonitoringItem[] = [
  { endpoint_id: 1, endpoint_name: 'Auth API', monitoring_started_at: '2025-01-12T08:00:00Z', monitoring_duration_days: 45 },
  { endpoint_id: 2, endpoint_name: 'Payments Gateway', monitoring_started_at: '2025-02-01T08:00:00Z', monitoring_duration_days: 30 },
]

export const mockDashboardIncidents: DashboardIncident[] = [
  { id: 1, endpoint_id: 2, started_at: '2026-06-25T14:22:00Z', resolved_at: null, is_resolved: false },
  { id: 2, endpoint_id: 4, started_at: '2026-06-24T09:15:00Z', resolved_at: '2026-06-24T10:45:00Z', is_resolved: true },
  { id: 3, endpoint_id: 7, started_at: '2026-06-23T18:30:00Z', resolved_at: '2026-06-23T19:00:00Z', is_resolved: true },
]

export const mockEndpoints: Endpoint[] = [
  { id: 1, name: 'Auth API', url: 'https://api.example.com/auth', expected_status: 200, status: 'healthy', last_checked: '2026-06-26T10:00:00Z', response_time: 124 },
  { id: 2, name: 'Payments Gateway', url: 'https://api.example.com/payments', expected_status: 200, status: 'unhealthy', last_checked: '2026-06-26T10:00:00Z', response_time: 890 },
  { id: 3, name: 'User Service', url: 'https://api.example.com/users', expected_status: 200, status: 'healthy', last_checked: '2026-06-26T10:00:00Z', response_time: 98 },
  { id: 4, name: 'Search Index', url: 'https://api.example.com/search', expected_status: 200, status: 'degraded', last_checked: '2026-06-26T10:00:00Z', response_time: 456 },
  { id: 5, name: 'Notifications', url: 'https://api.example.com/notifications', expected_status: 200, status: 'healthy', last_checked: '2026-06-26T10:00:00Z', response_time: 156 },
  { id: 6, name: 'Analytics API', url: 'https://api.example.com/analytics', expected_status: 200, status: 'healthy', last_checked: '2026-06-26T09:58:00Z', response_time: 210 },
  { id: 7, name: 'Inventory Service', url: 'https://api.example.com/inventory', expected_status: 200, status: 'unhealthy', last_checked: '2026-06-26T09:55:00Z', response_time: 1200 },
  { id: 8, name: 'CDN Health', url: 'https://cdn.example.com/health', expected_status: 200, status: 'healthy', last_checked: '2026-06-26T09:59:00Z', response_time: 45 },
]

export const mockHealthChecks: HealthCheck[] = [
  { id: 1, endpoint_id: 1, endpoint_name: 'Auth API', status_code: 200, response_time: 124, success: true, checked_at: '2026-06-26T10:00:00Z' },
  { id: 2, endpoint_id: 2, endpoint_name: 'Payments Gateway', status_code: 503, response_time: 890, success: false, checked_at: '2026-06-26T10:00:00Z' },
  { id: 3, endpoint_id: 3, endpoint_name: 'User Service', status_code: 200, response_time: 98, success: true, checked_at: '2026-06-26T10:00:00Z' },
  { id: 4, endpoint_id: 4, endpoint_name: 'Search Index', status_code: 200, response_time: 456, success: true, checked_at: '2026-06-26T09:59:00Z' },
  { id: 5, endpoint_id: 5, endpoint_name: 'Notifications', status_code: 200, response_time: 156, success: true, checked_at: '2026-06-26T09:58:00Z' },
  { id: 6, endpoint_id: 7, endpoint_name: 'Inventory Service', status_code: 500, response_time: 1200, success: false, checked_at: '2026-06-26T09:55:00Z' },
]

export const mockIncidents: Incident[] = [
  { id: 1, endpoint_id: 2, endpoint_name: 'Payments Gateway', severity: 'critical', started_at: '2026-06-25T14:22:00Z', resolved_at: null, is_resolved: false, status: 'active', duration_minutes: 1180 },
  { id: 2, endpoint_id: 7, endpoint_name: 'Inventory Service', severity: 'high', started_at: '2026-06-25T08:00:00Z', resolved_at: null, is_resolved: false, status: 'active', duration_minutes: 1560 },
  { id: 3, endpoint_id: 4, endpoint_name: 'Search Index', severity: 'medium', started_at: '2026-06-24T09:15:00Z', resolved_at: '2026-06-24T10:45:00Z', is_resolved: true, status: 'resolved', duration_minutes: 90 },
  { id: 4, endpoint_id: 2, endpoint_name: 'Payments Gateway', severity: 'high', started_at: '2026-06-20T16:00:00Z', resolved_at: '2026-06-20T16:35:00Z', is_resolved: true, status: 'resolved', duration_minutes: 35 },
]

export const mockResponseTimeChart = [
  { time: '00:00', ms: 145 },
  { time: '04:00', ms: 132 },
  { time: '08:00', ms: 198 },
  { time: '12:00', ms: 245 },
  { time: '16:00', ms: 187 },
  { time: '20:00', ms: 156 },
  { time: 'Now', ms: 172 },
]

export const mockRequestVolumeChart = [
  { time: '00:00', requests: 4200 },
  { time: '04:00', requests: 2800 },
  { time: '08:00', requests: 8900 },
  { time: '12:00', requests: 12400 },
  { time: '16:00', requests: 9800 },
  { time: '20:00', requests: 6500 },
  { time: 'Now', requests: 7200 },
]
