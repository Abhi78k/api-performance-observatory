import { cn } from '@/utils/cn'
import type { ButtonHTMLAttributes, ReactNode } from 'react'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'contained' | 'outlined' | 'text' | 'gradient'
  color?: 'info' | 'primary' | 'success' | 'error' | 'white'
  size?: 'small' | 'medium' | 'large'
  fullWidth?: boolean
  loading?: boolean
  iconOnly?: boolean
  children: ReactNode
}

const colorClasses = {
  info: 'btn-gradient-info text-white hover:opacity-90',
  primary: 'btn-gradient-primary text-white hover:opacity-90',
  success: 'bg-success text-white hover:bg-success/90',
  error: 'bg-error text-white hover:bg-error/90',
  white: 'bg-white/10 text-white hover:bg-white/20 border border-border',
}

export function Button({
  variant = 'contained',
  color = 'info',
  size = 'medium',
  fullWidth,
  loading,
  iconOnly,
  className,
  children,
  disabled,
  ...props
}: ButtonProps) {
  return (
    <button
      className={cn(
        'inline-flex items-center justify-center gap-2 rounded-lg font-medium transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed',
        size === 'small' && (iconOnly ? 'h-8 w-8 p-0' : 'h-8 px-3 text-xs'),
        size === 'medium' && (iconOnly ? 'h-10 w-10 p-0' : 'h-10 px-4 text-sm'),
        size === 'large' && (iconOnly ? 'h-12 w-12 p-0' : 'h-12 px-6 text-base'),
        variant === 'contained' && colorClasses[color],
        variant === 'gradient' && colorClasses[color],
        variant === 'outlined' && 'border border-border text-text-focus bg-transparent hover:bg-white/5',
        variant === 'text' && 'text-info bg-transparent hover:bg-white/5',
        fullWidth && 'w-full',
        className,
      )}
      disabled={disabled || loading}
      {...props}
    >
      {loading ? (
        <span className="h-4 w-4 animate-spin rounded-full border-2 border-white/30 border-t-white" />
      ) : (
        children
      )}
    </button>
  )
}
