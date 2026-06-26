import { cn } from '@/utils/cn'
import type { InputHTMLAttributes, ReactNode } from 'react'

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  icon?: ReactNode
  iconPosition?: 'left' | 'right'
  error?: boolean
}

export function Input({ icon, iconPosition = 'left', error, className, ...props }: InputProps) {
  const inputClasses = cn(
    'w-full rounded-lg bg-input-bg px-4 py-2.5 text-sm text-text-focus placeholder:text-grey-500 outline-none transition-colors',
    'border border-border focus:border-info/60',
    Boolean(icon) && iconPosition === 'left' && 'pl-10',
    Boolean(icon) && iconPosition === 'right' && 'pr-10',
    error && 'border-error/60',
    className,
  )

  if (!icon) {
    return (
      <div className="input-gradient-border">
        <input className={inputClasses} {...props} />
      </div>
    )
  }

  return (
    <div className="input-gradient-border">
      <div className="relative rounded-lg bg-input-bg">
        {iconPosition === 'left' && (
          <span className="absolute left-3 top-1/2 -translate-y-1/2 text-grey-400">{icon}</span>
        )}
        <input className={inputClasses} {...props} />
        {iconPosition === 'right' && (
          <span className="absolute right-3 top-1/2 -translate-y-1/2 text-grey-400">{icon}</span>
        )}
      </div>
    </div>
  )
}
