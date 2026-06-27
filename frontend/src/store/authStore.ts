import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import type { User } from '@/types/api'

interface AuthState {
  user: User | null
  isAuthenticated: boolean
  isCheckingAuth: boolean
  setAuth: (user: User) => void
  setUser: (user: User) => void
  setCheckingAuth: (isCheckingAuth: boolean) => void
  logout: () => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      isAuthenticated: false,
      isCheckingAuth: true,
      setAuth: (user) => set({ user, isAuthenticated: true, isCheckingAuth: false }),
      setUser: (user) => set({ user }),
      setCheckingAuth: (isCheckingAuth) => set({ isCheckingAuth }),
      logout: () => set({ user: null, isAuthenticated: false, isCheckingAuth: false }),
    }),
    {
      name: 'apo-auth',
      partialize: (state) => ({
        user: state.user,
        isAuthenticated: state.isAuthenticated,
      }),
    },
  ),
)
