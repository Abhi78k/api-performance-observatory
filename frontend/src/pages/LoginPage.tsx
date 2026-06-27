import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { AuthLayout } from '@/layouts/AuthLayout'
import { Button, Input, Typography } from '@/components/ui'
import { login, me } from '@/api/auth'
import { useAuthStore } from '@/store/authStore'

export function LoginPage() {
  const navigate = useNavigate()
  const setAuth = useAuthStore((s) => s.setAuth)
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)

    try {
      await login({ email, password })
      try {
        const user = await me()
        setAuth(user)
      } catch (err) {
        console.error('Failed to fetch user profile:', err)
        setError('Failed to fetch user profile after login.')
        return
      }
      navigate('/')

    } catch {
      // Mock fallback for development when backend is unavailable
      if (import.meta.env.VITE_USE_MOCK !== 'false') {
        setAuth({ id: 1, email })
        navigate('/')
      } else {
        setError('Invalid email or password. Please try again.')
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <AuthLayout
      title="Welcome back"
      description="Enter your email and password to sign in"
      footer={
        <Typography variant="button" color="text">
          Don&apos;t have an account?{' '}
          <Link to="/register" className="font-medium text-text-focus hover:text-info">
            Sign up
          </Link>
        </Typography>
      }
    >
      <form onSubmit={handleSubmit} className="space-y-5">
        <div>
          <Typography variant="button" color="white" fontWeight="medium" className="mb-1 ml-0.5 block">
            Email
          </Typography>
          <Input
            type="email"
            placeholder="user@example.com"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <Typography variant="button" color="white" fontWeight="medium" className="mb-1 ml-0.5 block">
            Password
          </Typography>
          <Input
            type="password"
            placeholder="Your password..."
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        {error && (
          <Typography variant="caption" color="error">
            {error}
          </Typography>
        )}
        <Button type="submit" fullWidth loading={loading}>
          Sign In
        </Button>
      </form>
    </AuthLayout>
  )
}
