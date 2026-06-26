import { cn } from '@/utils/cn'
import type { ElementType, HTMLAttributes } from 'react'

type Variant = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'subtitle1' | 'subtitle2' | 'body1' | 'body2' | 'caption' | 'button' | 'lg'

interface TypographyProps extends HTMLAttributes<HTMLElement> {
  variant?: Variant
  component?: ElementType
  color?: 'white' | 'text' | 'info' | 'success' | 'warning' | 'error'
  fontWeight?: 'regular' | 'medium' | 'bold'
}

const variantClasses: Record<Variant, string> = {
  h1: 'text-4xl font-bold',
  h2: 'text-3xl font-bold',
  h3: 'text-2xl font-bold',
  h4: 'text-xl font-bold',
  h5: 'text-lg font-bold',
  h6: 'text-base font-bold',
  subtitle1: 'text-base font-medium',
  subtitle2: 'text-sm font-medium',
  body1: 'text-base',
  body2: 'text-sm',
  caption: 'text-xs',
  button: 'text-sm font-medium',
  lg: 'text-lg font-bold',
}

const variantTags: Record<Variant, ElementType> = {
  h1: 'h1',
  h2: 'h2',
  h3: 'h3',
  h4: 'h4',
  h5: 'h5',
  h6: 'h6',
  subtitle1: 'p',
  subtitle2: 'p',
  body1: 'p',
  body2: 'p',
  caption: 'span',
  button: 'span',
  lg: 'p',
}

const colorClasses = {
  white: 'text-text-focus',
  text: 'text-text',
  info: 'text-info',
  success: 'text-success',
  warning: 'text-warning',
  error: 'text-error',
}

const weightClasses = {
  regular: 'font-normal',
  medium: 'font-medium',
  bold: 'font-bold',
}

export function Typography({
  variant = 'body1',
  component,
  color = 'text',
  fontWeight,
  className,
  children,
  ...props
}: TypographyProps) {
  const Tag = component ?? variantTags[variant]
  return (
    <Tag
      className={cn(
        variantClasses[variant],
        colorClasses[color],
        fontWeight && weightClasses[fontWeight],
        className,
      )}
      {...props}
    >
      {children}
    </Tag>
  )
}
