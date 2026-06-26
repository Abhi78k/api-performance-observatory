import { Link } from 'react-router-dom'
import { Globe2 } from 'lucide-react'
import { Typography } from '@/components/ui'
import type { ReactNode } from 'react'

interface AuthLayoutProps {
  title: string
  description: string
  children: ReactNode
  footer?: ReactNode
}

export function AuthLayout({ title, description, children, footer }: AuthLayoutProps) {
  return (
    <div className="flex min-h-screen">
      <div className="hidden w-1/2 flex-col justify-between bg-gradient-to-br from-[#0f123b] via-[#090d2e] to-[#020515] p-12 lg:flex">
        <div className="flex items-center gap-3">
          <div className="flex h-10 w-10 items-center justify-center rounded-lg btn-gradient-info">
            <Globe2 className="h-6 w-6 text-white" />
          </div>
          <Typography variant="h6" color="white" fontWeight="bold">
            API Performance Observatory
          </Typography>
        </div>
        <div>
          <Typography variant="caption" color="text" className="uppercase tracking-widest">
            Global API Monitoring Platform
          </Typography>
          <Typography variant="h2" color="white" fontWeight="bold" className="mt-2 max-w-md">
            Monitor endpoint health, response times, and incidents worldwide
          </Typography>
        </div>
        <Typography variant="caption" color="text">
          © 2026 API Performance Observatory
        </Typography>
      </div>

      <div className="flex w-full flex-col justify-center px-6 py-12 lg:w-1/2 lg:px-16">
        <div className="mx-auto w-full max-w-md">
          <div className="mb-8 lg:hidden">
            <Link to="/" className="flex items-center gap-2">
              <div className="flex h-9 w-9 items-center justify-center rounded-lg btn-gradient-info">
                <Globe2 className="h-5 w-5 text-white" />
              </div>
              <Typography variant="button" color="white" fontWeight="bold">
                API Performance Observatory
              </Typography>
            </Link>
          </div>

          <Typography variant="h4" color="white" fontWeight="bold" className="mb-1">
            {title}
          </Typography>
          <Typography variant="body2" color="text" className="mb-8">
            {description}
          </Typography>

          {children}

          {footer && <div className="mt-6 text-center">{footer}</div>}
        </div>
      </div>
    </div>
  )
}
