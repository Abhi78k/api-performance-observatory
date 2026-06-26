import { AlertTriangle, RefreshCw } from 'lucide-react'
import { Button } from './Button'
import { Typography } from './Typography'

interface ErrorStateProps {
  title?: string
  message?: string
  onRetry?: () => void
}

export function ErrorState({
  title = 'Something went wrong',
  message = 'Failed to load data. Please try again.',
  onRetry,
}: ErrorStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-12 text-center">
      <div className="mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-error/10">
        <AlertTriangle className="h-8 w-8 text-error" />
      </div>
      <Typography variant="subtitle1" color="white" fontWeight="bold" className="mb-1">
        {title}
      </Typography>
      <Typography variant="body2" color="text" className="mb-4 max-w-sm">
        {message}
      </Typography>
      {onRetry && (
        <Button variant="outlined" color="white" onClick={onRetry}>
          <RefreshCw className="h-4 w-4" />
          Retry
        </Button>
      )}
    </div>
  )
}
