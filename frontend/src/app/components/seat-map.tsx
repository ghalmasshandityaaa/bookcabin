"use client"

import { useState, useEffect } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Loader2 } from "lucide-react"
import { apiClient } from "@/lib/api-client"

interface Seat {
  row_number: number
  seat: string
  assigned: boolean
  generated?: boolean // New property to mark generated seats
}

interface SeatMapProps {
  aircraftType: string
  generatedSeats: string[]
  onSeatsLoaded?: (seats: Seat[]) => void
}

export default function SeatMap({ aircraftType, generatedSeats, onSeatsLoaded }: SeatMapProps) {
  const [seats, setSeats] = useState<Seat[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (aircraftType) {
      fetchSeats()
    }
  }, [aircraftType])

  // Update seats when generatedSeats changes
  useEffect(() => {
    if (generatedSeats.length > 0 && seats.length > 0) {
      updateSeatsWithGenerated()
    }
  }, [generatedSeats])

  const fetchSeats = async () => {
    try {
      setLoading(true)
      setError(null)

      const response = await fetch(`http://localhost:4000/api/aircraft/seatsss?type=${encodeURIComponent(aircraftType)}`)

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const contentType = response.headers.get("content-type")
      if (!contentType || !contentType.includes("application/json")) {
        throw new Error("Response is not JSON")
      }

      const result = await response.json()

      if (result.success) {
        const seatsData = result.data.map((seat: Seat) => ({
          ...seat,
          generated: false,
        }))
        setSeats(seatsData)
        onSeatsLoaded?.(seatsData)
      } else {
        throw new Error(result.error || "Invalid response format")
      }
    } catch (err) {
      console.error("Error fetching seats:", err)
      setError(err instanceof Error ? err.message : "Failed to load seats")
    } finally {
      setLoading(false)
    }
  }

  const updateSeatsWithGenerated = () => {
    setSeats((prevSeats) =>
      prevSeats.map((seat) => {
        const seatCode = `${seat.row_number}${seat.seat}`
        return {
          ...seat,
          generated: generatedSeats.includes(seatCode),
        }
      }),
    )
  }

  const getSeatColor = (seat: Seat) => {
    if (seat.generated) {
      return "bg-blue-500 text-white border-blue-600" // Generated seats (blue)
    } else if (seat.assigned) {
      return "bg-green-500 text-white border-green-600" // Previously assigned seats (green)
    } else {
      return "bg-gray-100 text-gray-700 border-gray-300 hover:bg-gray-200" // Available seats
    }
  }

  const getSeatStatus = (seat: Seat) => {
    if (seat.generated) {
      return "Generated"
    } else if (seat.assigned) {
      return "Assigned"
    } else {
      return "Available"
    }
  }

  // Group seats by row
  const seatsByRow = seats.reduce(
    (acc, seat) => {
      if (!acc[seat.row_number]) {
        acc[seat.row_number] = []
      }
      acc[seat.row_number].push(seat)
      return acc
    },
    {} as Record<number, Seat[]>,
  )

  // Sort rows numerically
  const sortedRows = Object.keys(seatsByRow)
    .map(Number)
    .sort((a, b) => a - b)

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
      {/* Legend - Centered */}
      <div className="flex flex-wrap gap-4 justify-center">
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

      {/* Aircraft Type Badge - Centered */}
      <div className="text-center">
        <Badge variant="outline" className="text-lg px-4 py-2">
          {aircraftType}
        </Badge>
      </div>

      {/* Seat Map - Centered Layout */}
      <div className="max-h-96 overflow-y-auto">
        <div className="space-y-2 flex flex-col items-center">
          {sortedRows.map((rowNumber) => {
            const rowSeats = seatsByRow[rowNumber].sort((a, b) => a.seat.localeCompare(b.seat))

            return (
              <div key={rowNumber} className="flex items-center gap-2">
                {/* Row Number */}
                <div className="w-8 text-center font-mono text-sm font-semibold text-gray-600">{rowNumber}</div>

                {/* Seats in Row - Centered */}
                <div className="flex gap-1 justify-center">
                  {rowSeats.map((seat) => {
                    const seatCode = `${seat.row_number}${seat.seat}`
                    return (
                      <div
                        key={seatCode}
                        className={`
                          w-10 h-10 rounded border-2 flex items-center justify-center
                          text-xs font-semibold transition-colors cursor-default
                          ${getSeatColor(seat)}
                        `}
                        title={`Seat ${seatCode} - ${getSeatStatus(seat)}`}
                      >
                        {seat.seat}
                      </div>
                    )
                  })}
                </div>
              </div>
            )
          })}
        </div>
      </div>

      {/* Generated Seats Summary - Centered with Smaller Boxes */}
      {generatedSeats.length > 0 && (
        <Card className="bg-blue-50 border-blue-200">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm text-blue-800 text-center">Generated Voucher Seats</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2 justify-center">
              {generatedSeats.map((seat) => (
                <div
                  key={seat}
                  className="bg-blue-500 text-white px-3 py-2 rounded-md font-semibold text-sm border-2 border-blue-600 min-w-[2.5rem] text-center"
                >
                  {seat}
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  )
}
