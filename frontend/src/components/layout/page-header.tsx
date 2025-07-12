"use client"

import { Plane } from "lucide-react"

export function PageHeader() {
  return (
    <div className="text-center py-6">
      <div className="flex items-center justify-center gap-3 mb-4">
        <Plane className="h-8 w-8 text-blue-600" />
        <h1 className="text-3xl font-bold text-gray-900">Airline Voucher Assignment</h1>
      </div>
      <p className="text-gray-600">Generate random seat assignments for airline campaign vouchers</p>
    </div>
  )
}
