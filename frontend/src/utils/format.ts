export function formatMs(ms: number): string {
  if (ms < 1000) return `${Math.round(ms)}ms`;
  return `${(ms / 1000).toFixed(2)}s`;
}

export function clampPercentage(val: number): number {
  if (isNaN(val) || !isFinite(val)) return 0;
  return Math.min(100, Math.max(0, val));
}

export function safeNumber(val: any, defaultVal = 0): number {
  const num = Number(val);
  if (isNaN(num) || !isFinite(num) || num < 0) return defaultVal;
  return num;
}

export function safeInteger(val: any): number {
  return Math.round(safeNumber(val));
}

export function formatLatency(ms: number): string {
  const safeMs = safeNumber(ms);
  return formatMs(safeMs);
}

export function formatPercent(value: number, decimals = 1): string {
  const clamped = clampPercentage(value);
  return `${clamped.toFixed(decimals)}%`;
}

export function formatDate(iso: string): string {
  return new Date(iso).toLocaleString(undefined, {
    month: "short",
    day: "numeric",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

export function formatRelativeTime(iso: string): string {
  const diff = Date.now() - new Date(iso).getTime();
  const minutes = Math.floor(diff / 60000);
  if (minutes < 1) return "Just now";
  if (minutes < 60) return `${minutes}m ago`;
  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours}h ago`;
  const days = Math.floor(hours / 24);
  return `${days}d ago`;
}

export function formatDuration(minutes: number): string {
  if (minutes < 60) return `${Math.trunc(minutes)}m`;
  const hours = Math.floor(minutes / 60);
  const mins = minutes % 60;
  if (hours < 24)
    return mins > 0 ? `${hours}h ${Math.trunc(mins)}m` : `${hours}h`;
  const days = Math.floor(hours / 24);
  const remHours = hours % 24;
  return remHours > 0 ? `${days}d ${remHours}h` : `${days}d`;
}

export function getStatusColor(
  status: string,
): "success" | "error" | "warning" | "info" {
  const s = status.toLowerCase();
  if (s === "healthy" || s === "resolved" || s === "success") return "success";
  if (s === "unhealthy" || s === "failed" || s === "critical" || s === "active")
    return "error";
  if (s === "degraded" || s === "slow" || s === "medium" || s === "high")
    return "warning";
  return "info";
}
