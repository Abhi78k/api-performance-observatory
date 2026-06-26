import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { AuthLayout } from '@/layouts/AuthLayout'
import { Button, Input, Typography } from '@/components/ui'
import { register } from '@/api/auth'

export function RegisterPage() {
  const navigate = useNavigate()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')

    if (password !== confirmPassword) {
      setError('Passwords do not match')
      return
    }

    setLoading(true)
    try {
      // TODO(API): Connect register form to POST /auth/register — wired below
      await register({ email, password })
      navigate('/login', { state: { message: 'Registration successful. Please sign in.' } })
    } catch {
      if (import.meta.env.VITE_USE_MOCK !== 'false') {
        navigate('/login', { state: { message: 'Registration successful (mock). Please sign in.' } })
      } else {
        setError('Registration failed. Please try again.')
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <AuthLayout
      title="Create an account"
      description="Register to start monitoring your APIs"
      footer={
        <Typography variant="button" color="text">
          Already have an account?{' '}
          <Link to="/login" className="font-medium text-text-focus hover:text-info">
            Sign in
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
            placeholder="Minimum 8 characters"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            minLength={8}
          />
        </div>
        <div>
          <Typography variant="button" color="white" fontWeight="medium" className="mb-1 ml-0.5 block">
            Confirm Password
          </Typography>
          <Input
            type="password"
            placeholder="Confirm your password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            required
          />
        </div>
        {error && (
          <Typography variant="caption" color="error">
            {error}
          </Typography>
        )}
        <Button type="submit" fullWidth loading={loading}>
          Create Account
        </Button>
      </form>
    </AuthLayout>
  )
}
