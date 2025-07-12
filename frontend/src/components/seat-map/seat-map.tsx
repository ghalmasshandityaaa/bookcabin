"use client"

import { useEffect } from "react"
import { Badge } from "@/components/ui/badge"
import { Loader2, Plane } from "lucide-react"
import { useSeatManagement } from "@/hooks/use-seat-management"
import { SeatLegend } from "./seat-legend"
import { SeatGrid } from "./seat-grid"
import { GeneratedSeatsSummary } from "./generated-seats-summary"

interface SeatMapProps {
  aircraftType: string
  currentGeneratedSeats: string[]
  allGeneratedSeats: string[]
}

export function SeatMap({ aircraftType, currentGeneratedSeats, allGeneratedSeats }: SeatMapProps) {
  const { seats, loading, error, fetchSeats } = useSeatManagement()

  useEffect(() => {
    if (aircraftType) {
      fetchSeats(aircraftType)
    }
  }, [aircraftType])

  if (!aircraftType) {
    return (
      <div className="flex items-center justify-center h-64 text-gray-500">
        <div className="text-center">
          <Plane className="h-12 w-12 mx-auto mb-4 opacity-50" />
          <p>Select aircraft type to view seat map</p>
        </div>
      </div>
    )
  }

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <Loader2 className="h-8 w-8 animate-spin text-blue-600" />
      </div>
    )
  }

  if (error) {
    return (
      <div className="text-center text-red-600 p-4">
        <p>Error loading seat map: {error}</p>
      </div>
    )
  }

  return (
    <div className="space-y-4">
      <SeatLegend />

      {/* Aircraft Type Badge - Centered */}
      <div className="text-center">
        <Badge variant="outline" className="text-lg px-4 py-2">
          {aircraftType}
        </Badge>
      </div>

      <SeatGrid 
        seats={seats} 
        currentGeneratedSeats={currentGeneratedSeats} 
        allGeneratedSeats={allGeneratedSeats} 
      />

      <GeneratedSeatsSummary seats={currentGeneratedSeats} />
    </div>
  )
}