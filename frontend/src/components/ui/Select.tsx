import { cn } from '@/utils/cn'
import type { SelectHTMLAttributes } from 'react'

interface SelectProps extends SelectHTMLAttributes<HTMLSelectElement> {
  label?: string
  options: { value: string; label: string }[]
}

export function Select({ label, options, className, ...props }: SelectProps) {
  return (
    <div>
      {label && (
        <label className="mb-1 ml-0.5 block text-sm font-medium text-text-focus">{label}</label>
      )}
      <div className="input-gradient-border">
        <select
          className={cn(
            'w-full rounded-lg bg-input-bg px-4 py-2.5 text-sm text-text-focus outline-none border border-border focus:border-info/60',
            className,
          )}
          {...props}
        >
          {options.map((opt) => (
            <option key={opt.value} value={opt.value}>
              {opt.label}
            </option>
          ))}
        </select>
      </div>
    </div>
  )
}
