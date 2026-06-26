import {
  Area,
  AreaChart,
  Bar,
  BarChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts'
import { Card, Typography } from '@/components/ui'

interface ChartCardProps {
  title: string
  subtitle?: string
  data: { time: string; ms?: number; requests?: number }[]
  dataKey: 'ms' | 'requests'
  color?: string
  unit?: string
  type?: 'area' | 'bar'
}

export function ChartCard({
  title,
  subtitle,
  data,
  dataKey,
  color = '#0075FF',
  unit = '',
  type = 'area',
}: ChartCardProps) {
  return (
    <Card>
      <Typography variant="lg" color="white" fontWeight="bold" className="mb-1">
        {title}
      </Typography>
      {subtitle && (
        <Typography variant="caption" color="text" className="mb-4 block">
          {subtitle}
        </Typography>
      )}
      <div className="h-48">
        <ResponsiveContainer width="100%" height="100%">
          {type === 'area' ? (
            <AreaChart data={data}>
              <defs>
                <linearGradient id={`gradient-${dataKey}`} x1="0" y1="0" x2="0" y2="1">
                  <stop offset="0%" stopColor={color} stopOpacity={0.4} />
                  <stop offset="100%" stopColor={color} stopOpacity={0} />
                </linearGradient>
              </defs>
              <CartesianGrid strokeDasharray="3 3" stroke="rgba(255,255,255,0.05)" />
              <XAxis dataKey="time" tick={{ fill: '#a0aec0', fontSize: 11 }} axisLine={false} tickLine={false} />
              <YAxis tick={{ fill: '#a0aec0', fontSize: 11 }} axisLine={false} tickLine={false} />
              <Tooltip
                contentStyle={{
                  background: 'rgba(6, 11, 40, 0.95)',
                  border: '1px solid rgba(226,232,240,0.2)',
                  borderRadius: '8px',
                  color: '#fff',
                }}
                formatter={(value) => [`${value}${unit}`, dataKey === 'ms' ? 'Response Time' : 'Requests']}
              />
              <Area
                type="monotone"
                dataKey={dataKey}
                stroke={color}
                fill={`url(#gradient-${dataKey})`}
                strokeWidth={2}
              />
            </AreaChart>
          ) : (
            <BarChart data={data}>
              <CartesianGrid strokeDasharray="3 3" stroke="rgba(255,255,255,0.05)" />
              <XAxis dataKey="time" tick={{ fill: '#a0aec0', fontSize: 11 }} axisLine={false} tickLine={false} />
              <YAxis tick={{ fill: '#a0aec0', fontSize: 11 }} axisLine={false} tickLine={false} />
              <Tooltip
                contentStyle={{
                  background: 'rgba(6, 11, 40, 0.95)',
                  border: '1px solid rgba(226,232,240,0.2)',
                  borderRadius: '8px',
                  color: '#fff',
                }}
              />
              <Bar dataKey={dataKey} fill={color} radius={[4, 4, 0, 0]} />
            </BarChart>
          )}
        </ResponsiveContainer>
      </div>
    </Card>
  )
}
