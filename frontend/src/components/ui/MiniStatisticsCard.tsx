import { Card } from './Card'
import { Typography } from './Typography'
import type { LucideIcon } from 'lucide-react'

interface MiniStatisticsCardProps {
  title: string
  value: string | number
  subtitle?: string
  subtitleColor?: 'success' | 'error' | 'warning' | 'info'
  icon: LucideIcon
  iconColor?: string
}

const subtitleColorMap = {
  success: 'text-success',
  error: 'text-error',
  warning: 'text-warning',
  info: 'text-info',
}

export function MiniStatisticsCard({
  title,
  value,
  subtitle,
  subtitleColor = 'success',
  icon: Icon,
  iconColor = '#0075FF',
}: MiniStatisticsCardProps) {
  return (
    <Card className="!p-[17px]">
      <div className="flex items-center justify-between gap-3">
        <div className="min-w-0 flex-1">
          <Typography variant="caption" color="text" className="capitalize opacity-70">
            {title}
          </Typography>
          <div className="mt-1 flex items-baseline gap-2">
            <Typography variant="subtitle1" color="white" fontWeight="bold">
              {value}
            </Typography>
            {subtitle && (
              <Typography variant="button" className={subtitleColorMap[subtitleColor]} fontWeight="bold">
                {subtitle}
              </Typography>
            )}
          </div>
        </div>
        <div
          className="flex h-12 w-12 shrink-0 items-center justify-center rounded-lg shadow-md"
          style={{ backgroundColor: iconColor }}
        >
          <Icon className="h-5 w-5 text-white" />
        </div>
      </div>
    </Card>
  )
}
