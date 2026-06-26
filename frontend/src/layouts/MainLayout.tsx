import { Outlet } from 'react-router-dom'
import { TopNav } from './TopNav'

export function MainLayout() {
  return (
    <div className="min-h-screen">
      <TopNav />
      <main className="mx-auto max-w-[1600px] px-4 py-6 lg:px-6">
        <Outlet />
      </main>
    </div>
  )
}
