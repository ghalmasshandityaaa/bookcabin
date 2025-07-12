"use client"

import type { Seat } from "@/types"

interface SeatGridProps {
  seats: Seat[]
  currentGeneratedSeats: string[] // Current generated seats (blue)
  allGeneratedSeats: string[] // All generated seats (for checking if previously generated)
}

export function SeatGrid({ seats, currentGeneratedSeats, allGeneratedSeats }: SeatGridProps) {
  const getSeatColor = (seat: Seat) => {
    const seatCode = `${seat.row_number}${seat.seat}`

    // Priority: Current generated seats (blue) > Previously generated (green) > Available (white)
    if (currentGeneratedSeats.includes(seatCode)) {
      return "bg-blue-500 text-white border-blue-600" // Current generated seats (blue)
    } else if (allGeneratedSeats.includes(seatCode) || seat.assigned) {
      return "bg-green-500 text-white border-green-600" // Previously generated/assigned seats (green)
    } else {
      return "bg-gray-100 text-gray-700 border-gray-300 hover:bg-gray-200" // Available seats
    }
  }

  const getSeatStatus = (seat: Seat) => {
    const seatCode = `${seat.row_number}${seat.seat}`

    if (currentGeneratedSeats.includes(seatCode)) {
      return "Just Generated"
    } else if (allGeneratedSeats.includes(seatCode) || seat.assigned) {
      return "Previously Assigned"
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

  return (
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
  )
}