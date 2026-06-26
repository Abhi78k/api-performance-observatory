import { cn } from '@/utils/cn'
import type { ReactNode } from 'react'

type BadgeColor = 'info' | 'success' | 'warning' | 'error' | 'primary' | 'secondary'

interface BadgeProps {
  color?: BadgeColor
  children: ReactNode
  className?: string
}

const colorClasses: Record<BadgeColor, string> = {
  info: 'bg-info/20 text-info border-info/30',
  success: 'bg-success/20 text-success border-success/30',
  warning: 'bg-warning/20 text-warning border-warning/30',
  error: 'bg-error/20 text-error border-error/30',
  primary: 'bg-primary/20 text-primary-focus border-primary/30',
  secondary: 'bg-secondary text-text border-border',
}

export function Badge({ color = 'info', children, className }: BadgeProps) {
  return (
    <span
      className={cn(
        'inline-flex items-center rounded-md border px-2 py-0.5 text-xs font-medium capitalize',
        colorClasses[color],
        className,
      )}
    >
      {children}
    </span>
  )
}
