import { cn } from '@/utils/cn'
import type { HTMLAttributes, ReactNode } from 'react'

interface CardProps extends HTMLAttributes<HTMLDivElement> {
  children: ReactNode
  padding?: boolean
}

export function Card({ children, padding = true, className, ...props }: CardProps) {
  return (
    <div className={cn('card-gradient rounded-xl', padding && 'p-5', className)} {...props}>
      {children}
    </div>
  )
}
