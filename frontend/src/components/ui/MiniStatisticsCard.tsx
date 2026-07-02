import { Card } from "./Card";
import { Typography } from "./Typography";
import type { LucideIcon } from "lucide-react";

interface MiniStatisticsCardProps {
  title: string;
  value: string | number;
  subtitle?: string;
  subtitleColor?: "success" | "error" | "warning" | "info";
  icon: LucideIcon;
}

const subtitleColorMap = {
  success: "text-success bg-success/10",
  error: "text-error bg-error/10",
  warning: "text-warning bg-warning/10",
  info: "text-info bg-info/10",
};

export function MiniStatisticsCard({
  title,
  value,
  subtitle,
  subtitleColor = "success",
  icon: Icon,
}: MiniStatisticsCardProps) {
  // Define your exact dark blue color here
  const darkBlueColor = "#0075FF";

  return (
    <Card className="!p-[17px] bg-slate-900/50 backdrop-blur-md border border-slate-800">
      <div className="flex items-center justify-between gap-4">
        <div className="min-w-0 flex-1">
          {/* Muted, tracked-out upper/capitalize title */}
          <Typography
            variant="caption"
            color="text"
            className="uppercase tracking-wider text-xs font-semibold text-slate-400"
          >
            {title}
          </Typography>

          <div className="mt-2 flex items-baseline gap-2.5">
            <Typography
              variant="subtitle1"
              color="white"
              className="text-2xl font-bold tracking-tight"
            >
              {value}
            </Typography>
            {subtitle && (
              /* Soft badge styling for the percentage instead of just raw text */
              <span
                className={`inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium ${subtitleColorMap[subtitleColor]}`}
              >
                {subtitle}
              </span>
            )}
          </div>
        </div>

        {/* Unified dark blue glass-tinted icon container */}
        <div
          className="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl border border-white/5 shadow-inner"
          style={{
            backgroundColor: `${darkBlueColor}15`, // Appends 15 for ~8% opacity tint
          }}
        >
          <Icon className="h-5 w-5" style={{ color: darkBlueColor }} />
        </div>
      </div>
    </Card>
  );
}
