# API Performance Observatory — Frontend

Enterprise API monitoring dashboard built with React + Vite, matching the Vision UI design system from the parent repository.

## Tech Stack

- React 19 + TypeScript
- Vite 8
- Tailwind CSS 4
- React Router 7
- TanStack Query
- Axios
- Zustand (auth only)
- Recharts
- Cobe.js (globe)
- Lucide Icons

## Setup

```bash
cd frontend
npm install
cp .env.example .env
npm run dev
```

Open [http://localhost:5173](http://localhost:5173)

## Backend Integration

The app proxies API requests to `http://localhost:8080` via `/api` (see `vite.config.ts`).

Set `VITE_USE_MOCK=false` in `.env` to disable mock fallbacks when the backend is running.

### API Endpoints

| Screen | Endpoints |
|--------|-----------|
| Login | `POST /auth/login` |
| Register | `POST /auth/register` |
| Profile | `GET /auth/me` |
| Dashboard | `GET /dashboard/*` |
| Endpoints | `GET/POST/PUT/DELETE /endpoints` |
| Endpoint Details | `GET /endpoints/{id}`, `/stats`, `/monitoring`, `/incidents` |
| Health Checks | `GET /healthchecks` |
| Incidents | `GET /incidents`, `/incidents/active` |

## Project Structure

```
src/
├── api/           # Axios API layer
├── components/    # Reusable UI (Vision UI design tokens)
├── features/      # Feature-specific components
├── hooks/         # TanStack Query hooks
├── layouts/       # TopNav, AuthLayout, MainLayout
├── mocks/         # Mock data with TODO(API) markers
├── pages/         # Route pages
├── routes/        # Router + protected routes
├── store/         # Zustand auth store
├── types/         # TypeScript types
└── utils/         # Helpers
```

## Scripts

| Command | Description |
|---------|-------------|
| `npm run dev` | Start dev server |
| `npm run build` | Production build |
| `npm run preview` | Preview production build |

## Authentication

Uses JWT Bearer tokens stored in Zustand (persisted). Unauthenticated users are redirected to `/login`.

For local development without a backend, any credentials will authenticate via mock fallback when `VITE_USE_MOCK=true`.
