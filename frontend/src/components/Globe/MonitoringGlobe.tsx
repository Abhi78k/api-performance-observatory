import { useEffect, useRef } from 'react'
import createGlobe from 'cobe'
import { MONITORING_NODES, MOCK_GLOBE_ARCS } from '@/mocks/data'
import type { ArcType, GlobeArc } from '@/types/api'

const ARC_COLORS: Record<ArcType, [number, number, number]> = {
  active: [0.0, 0.46, 1.0],
  success: [0.004, 0.71, 0.45],
  slow: [1.0, 0.71, 0.28],
  failed: [0.89, 0.1, 0.1],
}

interface MonitoringGlobeProps {
  arcs?: GlobeArc[]
  className?: string
}

export function MonitoringGlobe({ arcs = MOCK_GLOBE_ARCS, className }: MonitoringGlobeProps) {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const phiRef = useRef(0)
  const frameRef = useRef<number>(0)

  useEffect(() => {
    const canvas = canvasRef.current
    if (!canvas) return

    let width = 0
    let height = 0

    const onResize = () => {
      if (!canvas.parentElement) return
      width = canvas.parentElement.offsetWidth
      height = canvas.parentElement.offsetHeight
      canvas.width = width * 2
      canvas.height = height * 2
      canvas.style.width = `${width}px`
      canvas.style.height = `${height}px`
    }

    window.addEventListener('resize', onResize)
    onResize()

    const globe = createGlobe(canvas, {
      devicePixelRatio: 2,
      width: width * 2,
      height: height * 2,
      phi: 0,
      theta: 0.25,
      dark: 1,
      diffuse: 1.2,
      mapSamples: 16000,
      mapBrightness: 6,
      baseColor: [0.08, 0.12, 0.28],
      markerColor: [0.0, 0.46, 1.0],
      glowColor: [0.05, 0.08, 0.2],
      arcWidth: 0.4,
      arcHeight: 0.25,
      markerElevation: 0.02,
      markers: MONITORING_NODES.map((node) => ({
        location: [node.lat, node.lng],
        size: 0.06,
      })),
      arcs: arcs.map((arc) => ({
        from: [arc.start.lat, arc.start.lng] as [number, number],
        to: [arc.end.lat, arc.end.lng] as [number, number],
        color: ARC_COLORS[arc.type],
      })),
    })

    const animate = () => {
      phiRef.current += 0.003
      globe.update({
        phi: phiRef.current,
        width: width * 2,
        height: height * 2,
      })
      frameRef.current = requestAnimationFrame(animate)
    }

    frameRef.current = requestAnimationFrame(animate)

    return () => {
      globe.destroy()
      cancelAnimationFrame(frameRef.current)
      window.removeEventListener('resize', onResize)
    }
  }, [arcs])

  return (
    <div className={`relative ${className ?? ''}`}>
      <canvas ref={canvasRef} className="h-full w-full" />
      <div className="absolute bottom-2 left-1/2 flex -translate-x-1/2 gap-4 rounded-lg bg-black/40 px-4 py-2 text-xs backdrop-blur-sm">
        {(['active', 'success', 'slow', 'failed'] as ArcType[]).map((type) => {
          const [r, g, b] = ARC_COLORS[type]
          return (
            <div key={type} className="flex items-center gap-1.5 capitalize text-text">
              <span
                className="h-2 w-2 rounded-full"
                style={{ backgroundColor: `rgb(${r * 255}, ${g * 255}, ${b * 255})` }}
              />
              {type}
            </div>
          )
        })}
      </div>
    </div>
  )
}
