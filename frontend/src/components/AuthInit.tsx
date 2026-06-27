import { useEffect } from 'react'
import { useAuthStore } from '@/store/authStore'
import { me } from '@/api/auth'

interface AuthInitProps {
  children: React.ReactNode
}

export function AuthInit({ children }: AuthInitProps) {
  const { setAuth, logout, setCheckingAuth, isCheckingAuth } = useAuthStore()

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const user = await me()
        setAuth(user)
      } catch (err) {
        logout()
      } finally {
        setCheckingAuth(false)
      }
    }

    checkAuth()
  }, [setAuth, logout, setCheckingAuth])

  if (isCheckingAuth) {
    return (
      <div className="flex h-screen w-screen flex-col items-center justify-center bg-[#0a0f1d] text-white">
        <div className="relative flex items-center justify-center">
          <div className="h-16 w-16 animate-spin rounded-full border-4 border-indigo-500/20 border-t-indigo-500"></div>
          <div className="absolute h-10 w-10 animate-ping rounded-full bg-indigo-500/10"></div>
        </div>
        <span className="mt-4 text-sm font-medium text-slate-400 tracking-wider animate-pulse">Verifying Session...</span>
      </div>
    )
  }

  return <>{children}</>
}
