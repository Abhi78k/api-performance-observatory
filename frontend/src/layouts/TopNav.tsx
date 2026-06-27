import { Link, NavLink, useNavigate } from "react-router-dom";
import {
  Activity,
  AlertTriangle,
  Globe2,
  HeartPulse,
  LayoutDashboard,
  LogOut,
  Server,
  User,
} from "lucide-react";
import { useAuthStore } from "@/store/authStore";
import { Typography } from "@/components/ui";
import { cn } from "@/utils/cn";
import { logout as apiLogout } from "@/api/auth";

const navItems = [
  { to: "/", label: "Dashboard", icon: LayoutDashboard },
  { to: "/endpoints", label: "Endpoints", icon: Server },
  { to: "/health-checks", label: "Health Checks", icon: HeartPulse },
  { to: "/incidents", label: "Incidents", icon: AlertTriangle },
];

export function TopNav() {
  const navigate = useNavigate();
  const { user, logout } = useAuthStore();

  const handleLogout = async () => {
    try {
      await apiLogout();
    } catch (err) {
      console.error("Logout from server failed:", err);
    } finally {
      logout();
      navigate("/login");
    }
  };

  return (
    <header className="navbar-gradient sticky top-0 z-40 w-full">
      <div className="mx-auto flex h-16 max-w-[1600px] items-center justify-between px-4 lg:px-6">
        <div className="flex items-center gap-8">
          <Link to="/" className="flex items-center gap-2">
            <div className="flex h-9 w-9 items-center justify-center rounded-lg btn-gradient-info">
              <Globe2 className="h-5 w-5 text-white" />
            </div>
            <div className="hidden sm:block">
              <Typography
                variant="button"
                color="white"
                fontWeight="bold"
                className="leading-tight"
              >
                API Performance Observatory
              </Typography>
              {/*<Typography variant="caption" color="text" className="leading-tight">
                Observatory
              </Typography>*/}
            </div>
          </Link>

          <nav className="hidden items-center gap-1 md:flex">
            {navItems.map(({ to, label, icon: Icon }) => (
              <NavLink
                key={to}
                to={to}
                end={to === "/"}
                className={({ isActive }) =>
                  cn(
                    "flex items-center gap-2 rounded-lg px-3 py-2 text-sm font-medium transition-colors",
                    isActive
                      ? "bg-info/20 text-text-focus"
                      : "text-text hover:bg-white/5 hover:text-text-focus",
                  )
                }
              >
                <Icon className="h-4 w-4" />
                {label}
              </NavLink>
            ))}
          </nav>
        </div>

        <div className="flex items-center gap-3">
          <div className="hidden items-center gap-2 sm:flex">
            <Activity className="h-4 w-4 text-success animate-pulse-glow" />
            <Typography variant="caption" color="text">
              Live Monitoring
            </Typography>
          </div>
          <Link
            to="/profile"
            className="flex items-center gap-2 rounded-lg px-3 py-2 text-sm text-text transition-colors hover:bg-white/5 hover:text-text-focus"
          >
            <User className="h-4 w-4" />
            <span className="hidden sm:inline">{user?.email ?? "Profile"}</span>
          </Link>
          <button
            onClick={handleLogout}
            className="flex items-center gap-2 rounded-lg px-3 py-2 text-sm text-text transition-colors hover:bg-white/5 hover:text-error"
          >
            <LogOut className="h-4 w-4" />
            <span className="hidden sm:inline">Logout</span>
          </button>
        </div>
      </div>

      <nav className="flex gap-1 overflow-x-auto border-t border-border/50 px-4 py-2 md:hidden">
        {navItems.map(({ to, label, icon: Icon }) => (
          <NavLink
            key={to}
            to={to}
            end={to === "/"}
            className={({ isActive }) =>
              cn(
                "flex shrink-0 items-center gap-1.5 rounded-lg px-3 py-1.5 text-xs font-medium",
                isActive ? "bg-info/20 text-text-focus" : "text-text",
              )
            }
          >
            <Icon className="h-3.5 w-3.5" />
            {label}
          </NavLink>
        ))}
      </nav>
    </header>
  );
}
