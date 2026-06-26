import { Inbox } from 'lucide-react'
import { Typography } from './Typography'

interface EmptyStateProps {
  title?: string
  description?: string
  action?: React.ReactNode
}

export function EmptyState({
  title = 'No data found',
  description = 'There is nothing to display at the moment.',
  action,
}: EmptyStateProps) {
  return (
    <div className="flex flex-col items-center justify-center py-12 text-center">
      <div className="mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-white/5">
        <Inbox className="h-8 w-8 text-grey-400" />
      </div>
      <Typography variant="subtitle1" color="white" fontWeight="bold" className="mb-1">
        {title}
      </Typography>
      <Typography variant="body2" color="text" className="mb-4 max-w-sm">
        {description}
      </Typography>
      {action}
    </div>
  )
}
