# API Performance Observatory

API Performance Observatory is a full-stack application designed for comprehensive monitoring of API endpoints. It provides real-time insights into API health, performance metrics, and incident management through an intuitive dashboard.

The backend is built with Go (Gin), providing a high-performance API for data collection and retrieval, while the frontend is a modern React/Vite application with a unique, observability-focused design system.

## Features

*   **Centralized Dashboard**: Get a bird's-eye view of your entire API ecosystem, including endpoint health, overall uptime, and average response times.
*   **Endpoint Management**: Easily add, edit, and remove API endpoints for monitoring.
*   **Automated Health Checks**: The system periodically performs health checks on your endpoints, verifying status codes and measuring response latency.
*   **Incident Tracking**: Automatically creates incidents when an endpoint fails and resolves them upon recovery. View active and historical incidents.
*   **Performance Analytics**: Dive deep into endpoint statistics with detailed metrics on success rates, latency, and request volume.
*   **Global Monitoring Visualization**: A dynamic 3D globe visualizes monitoring nodes and simulated traffic patterns across different geographical regions.
*   **Secure Authentication**: User registration and login are secured using JWT.
*   **RESTful API with Documentation**: A well-defined backend API with Swagger documentation for easy integration and exploration.

## Tech Stack

| Component | Technology |
| :--- | :--- |
| **Backend** | Go, Gin, GORM, PostgreSQL, JWT |
| **Frontend**| React (Vite), TypeScript, Tailwind CSS, TanStack Query, Zustand, Recharts, Cobe.js |
| **Database** | PostgreSQL |

## Project Architecture

The application is structured as a monorepo with two main components:
-   **`/backend`**: A Go application that serves a RESTful API. It handles user authentication, endpoint management, data collection via a scheduled job runner, and serves aggregated statistics.
-   **`/frontend`**: A React single-page application that consumes the backend API to provide a user-friendly monitoring dashboard.

The backend scheduler runs health checks every minute, recording the results and automatically managing the lifecycle of incidents. The frontend uses TanStack Query for efficient data fetching and caching, and Zustand for global state management (auth).

## Getting Started

Follow these instructions to set up and run the API Performance Observatory on your local machine.

### Prerequisites

*   Go (version 1.22 or later)
*   Node.js (version 20.x or later)
*   npm
*   PostgreSQL

### Backend Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Abhi78k/api-performance-observatory.git
    cd api-performance-observatory/backend
    ```

2.  **Set up the database:**
    *   Ensure your PostgreSQL server is running.
    *   Create a new database for the application.

3.  **Configure Environment Variables:**
    *   Copy the example environment file:
        ```sh
        cp .env.example .env
        ```
    *   Edit the `.env` file with your PostgreSQL database credentials and a secret for JWT:
        ```
        DB_HOST=localhost
        DB_PORT=5432
        DB_USER=your_postgres_user
        DB_PASSWORD=your_postgres_password
        DB_NAME=your_database_name
        JWT_SECRET=your_jwt_secret
        ```

4.  **Install dependencies and run the server:**
    ```sh
    go mod tidy
    go run ./cmd/server/main.go
    ```
    The backend server will start on `http://localhost:8080`. The application will automatically run database migrations on startup.

### Frontend Setup

1.  **Navigate to the frontend directory:**
    ```sh
    # From the repository root
    cd frontend
    ```

2.  **Install dependencies:**
    ```sh
    npm install
    ```

3.  **Configure Environment Variables:**
    *   Copy the example environment file. The default configuration proxies API requests to the backend at `http://localhost:8080`.
        ```sh
        cp .env.example .env
        ```
    *   The `VITE_USE_MOCK` flag can be set to `true` to run the frontend with mock data, even if the backend is not running.

4.  **Run the development server:**
    ```sh
    npm run dev
    ```
    The frontend application will be available at `http://localhost:5173`.

## API Documentation

The backend API includes Swagger documentation. Once the backend server is running, you can access the interactive API documentation at:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
