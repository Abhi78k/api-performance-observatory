import { useEffect } from 'react'
import { User, Mail, Shield } from 'lucide-react'
import { Card, CardSkeleton, ErrorState, Typography } from '@/components/ui'
import { me } from '@/api/auth'
import { useAuthStore } from '@/store/authStore'
import { useQuery } from '@tanstack/react-query'

export function ProfilePage() {
  const { user, setUser, token } = useAuthStore()

  const profile = useQuery({
    queryKey: ['auth', 'me'],
    queryFn: me,
    enabled: !!token,
    retry: false,
  })

  useEffect(() => {
    if (profile.data) setUser(profile.data)
  }, [profile.data, setUser])

  const displayUser = profile.data ?? user

  if (profile.isLoading && !displayUser) return <CardSkeleton />

  return (
    <div className="space-y-6">
      <div>
        <Typography variant="h5" color="white" fontWeight="bold">
          Profile
        </Typography>
        <Typography variant="body2" color="text">
          Manage your account information
        </Typography>
      </div>

      {profile.isError && !displayUser && (
        <ErrorState
          onRetry={() => profile.refetch()}
          message="Failed to load profile. Using cached data if available."
        />
      )}

      <Card>
        <div className="flex items-center gap-4 mb-6">
          <div className="flex h-16 w-16 items-center justify-center rounded-xl btn-gradient-info">
            <User className="h-8 w-8 text-white" />
          </div>
          <div>
            <Typography variant="h6" color="white" fontWeight="bold">
              {displayUser?.email ?? 'User'}
            </Typography>
            <Typography variant="caption" color="text">
              Account ID: {displayUser?.id ?? '—'}
            </Typography>
          </div>
        </div>

        <dl className="space-y-4">
          <div className="flex items-center gap-3 rounded-lg bg-white/5 px-4 py-3">
            <Mail className="h-5 w-5 text-info" />
            <div>
              <Typography variant="caption" color="text">Email Address</Typography>
              <Typography variant="button" color="white">{displayUser?.email ?? '—'}</Typography>
            </div>
          </div>
          <div className="flex items-center gap-3 rounded-lg bg-white/5 px-4 py-3">
            <Shield className="h-5 w-5 text-success" />
            <div>
              <Typography variant="caption" color="text">Authentication</Typography>
              <Typography variant="button" color="white">JWT Bearer Token</Typography>
            </div>
          </div>
        </dl>
      </Card>
    </div>
  )
}
