"use client"

interface SeatLegendProps {
  className?: string
}

export function SeatLegend({ className }: SeatLegendProps) {
  return (
    <div className={`flex flex-wrap gap-4 justify-center ${className}`}>
      <div className="flex items-center gap-2">
        <div className="w-4 h-4 bg-blue-500 rounded border"></div>
        <span className="text-sm">Generated</span>
      </div>
      <div className="flex items-center gap-2">
        <div className="w-4 h-4 bg-green-500 rounded border"></div>
        <span className="text-sm">Assigned</span>
      </div>
      <div className="flex items-center gap-2">
        <div className="w-4 h-4 bg-gray-100 rounded border border-gray-300"></div>
        <span className="text-sm">Available</span>
      </div>
    </div>
  )
}
