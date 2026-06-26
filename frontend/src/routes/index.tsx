import { createBrowserRouter, Navigate } from 'react-router-dom'
import { MainLayout } from '@/layouts/MainLayout'
import { ProtectedRoute, PublicRoute } from '@/routes/ProtectedRoute'
import { DashboardPage } from '@/pages/DashboardPage'
import { LoginPage } from '@/pages/LoginPage'
import { RegisterPage } from '@/pages/RegisterPage'
import { EndpointsPage } from '@/pages/EndpointsPage'
import { EndpointDetailsPage } from '@/pages/EndpointDetailsPage'
import { HealthChecksPage } from '@/pages/HealthChecksPage'
import { IncidentsPage } from '@/pages/IncidentsPage'
import { ProfilePage } from '@/pages/ProfilePage'

export const router = createBrowserRouter([
  {
    element: <ProtectedRoute />,
    children: [
      {
        element: <MainLayout />,
        children: [
          { index: true, element: <DashboardPage /> },
          { path: 'endpoints', element: <EndpointsPage /> },
          { path: 'endpoints/:id', element: <EndpointDetailsPage /> },
          { path: 'health-checks', element: <HealthChecksPage /> },
          { path: 'incidents', element: <IncidentsPage /> },
          { path: 'profile', element: <ProfilePage /> },
        ],
      },
    ],
  },
  {
    element: <PublicRoute />,
    children: [
      { path: 'login', element: <LoginPage /> },
      { path: 'register', element: <RegisterPage /> },
    ],
  },
  { path: '*', element: <Navigate to="/" replace /> },
])
