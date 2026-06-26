import { cn } from '@/utils/cn'
import type { HTMLAttributes } from 'react'

interface BoxProps extends HTMLAttributes<HTMLDivElement> {
  variant?: 'transparent' | 'card' | 'gradient'
}

export function Box({ variant = 'transparent', className, children, ...props }: BoxProps) {
  return (
    <div
      className={cn(
        variant === 'card' && 'card-gradient rounded-xl',
        variant === 'gradient' && 'card-gradient rounded-xl p-5',
        className,
      )}
      {...props}
    >
      {children}
    </div>
  )
}
