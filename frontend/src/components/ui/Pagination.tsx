import { Button } from './Button'
import { ChevronLeft, ChevronRight } from 'lucide-react'

interface PaginationProps {
  page: number
  totalPages: number
  onPageChange: (page: number) => void
  hasNext: boolean
  hasPrevious: boolean
}

export function Pagination({
  page,
  totalPages,
  onPageChange,
  hasNext,
  hasPrevious,
}: PaginationProps) {
  if (totalPages <= 1) return null

  // Generate page numbers
  const pages = []
  const maxVisiblePages = 5
  let startPage = Math.max(1, page - 2)
  const endPage = Math.min(totalPages, startPage + maxVisiblePages - 1)
  
  if (endPage - startPage + 1 < maxVisiblePages) {
    startPage = Math.max(1, endPage - maxVisiblePages + 1)
  }

  for (let i = startPage; i <= endPage; i++) {
    pages.push(i)
  }

  return (
    <div className="flex items-center justify-between border-t border-white/10 px-4 py-3 sm:px-6 mt-4">
      <div className="flex flex-1 justify-between sm:hidden">
        <Button
          variant="outlined"
          size="small"
          disabled={!hasPrevious}
          onClick={() => onPageChange(page - 1)}
        >
          Previous
        </Button>
        <Button
          variant="outlined"
          size="small"
          disabled={!hasNext}
          onClick={() => onPageChange(page + 1)}
        >
          Next
        </Button>
      </div>
      <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
        <div>
          <TypographyText>
            Page <span className="font-semibold text-white">{page}</span> of{' '}
            <span className="font-semibold text-white">{totalPages}</span>
          </TypographyText>
        </div>
        <div>
          <nav className="isolate inline-flex -space-x-px rounded-md shadow-sm gap-1.5" aria-label="Pagination">
            <Button
              variant="outlined"
              size="small"
              iconOnly
              disabled={!hasPrevious}
              onClick={() => onPageChange(page - 1)}
            >
              <ChevronLeft className="h-4 w-4" />
            </Button>
            
            {pages.map((p) => (
              <Button
                key={p}
                variant={p === page ? 'contained' : 'outlined'}
                color={p === page ? 'info' : 'white'}
                size="small"
                className="w-8 h-8 p-0"
                onClick={() => onPageChange(p)}
              >
                {p}
              </Button>
            ))}

            <Button
              variant="outlined"
              size="small"
              iconOnly
              disabled={!hasNext}
              onClick={() => onPageChange(page + 1)}
            >
              <ChevronRight className="h-4 w-4" />
            </Button>
          </nav>
        </div>
      </div>
    </div>
  )
}

// Simple typography text wrapper to avoid importing full Typography if unnecessary, or we can use regular html.
function TypographyText({ children }: { children: React.ReactNode }) {
  return (
    <p className="text-xs text-[#a0aec0] font-medium pt-1">
      {children}
    </p>
  )
}
