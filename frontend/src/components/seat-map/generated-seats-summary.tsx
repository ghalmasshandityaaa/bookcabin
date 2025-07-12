"use client"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

interface GeneratedSeatsSummaryProps {
  seats: string[]
}

export function GeneratedSeatsSummary({ seats }: GeneratedSeatsSummaryProps) {
  if (seats.length === 0) return null

  return (
    <Card className="bg-blue-50 border-blue-200">
      <CardHeader className="pb-2">
        <CardTitle className="text-sm text-blue-800 text-center">Generated Voucher Seats</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="flex flex-wrap gap-2 justify-center">
          {seats.map((seat) => (
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
  )
}
