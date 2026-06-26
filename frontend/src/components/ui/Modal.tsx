import { X } from 'lucide-react'
import { Button } from './Button'
import { Typography } from './Typography'
import type { ReactNode } from 'react'

interface ModalProps {
  open: boolean
  onClose: () => void
  title: string
  children: ReactNode
  footer?: ReactNode
}

export function Modal({ open, onClose, title, children, footer }: ModalProps) {
  if (!open) return null

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div className="absolute inset-0 bg-black/60 backdrop-blur-sm" onClick={onClose} />
      <div className="card-gradient relative z-10 w-full max-w-lg rounded-xl p-6 shadow-2xl">
        <div className="mb-4 flex items-center justify-between">
          <Typography variant="h6" color="white" fontWeight="bold">
            {title}
          </Typography>
          <Button variant="text" iconOnly size="small" onClick={onClose} aria-label="Close">
            <X className="h-4 w-4" />
          </Button>
        </div>
        <div className="mb-4">{children}</div>
        {footer && <div className="flex justify-end gap-3">{footer}</div>}
      </div>
    </div>
  )
}
