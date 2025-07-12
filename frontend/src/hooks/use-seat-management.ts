"use client"

import { useState, useEffect, useCallback } from "react"
import type { Seat } from "@/types"
import { apiClient } from "@/lib/api-client"

export const useSeatManagement = () => {
  const [seats, setSeats] = useState<Seat[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [currentGeneratedSeats, setCurrentGeneratedSeats] = useState<string[]>([]) // Latest generated seats
  const [allGeneratedSeats, setAllGeneratedSeats] = useState<string[]>([]) // All generated seats ever

  const fetchSeats = async (aircraftType: string) => {
    if (!aircraftType) return

    try {
      setLoading(true)
      setError(null)

      const result = await apiClient.getAircraftSeats(aircraftType)

      if (result.success) {
        const seatsData = result.data.map((seat) => ({
          ...seat,
          generated: false,
        }))
        setSeats(seatsData)
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

  // Add new generated seats while keeping track of all previously generated seats
  const addGeneratedSeats = (newSeats: string[]) => {
    setCurrentGeneratedSeats(newSeats)
    setAllGeneratedSeats((prev) => {
      // Combine previous and new seats, removing duplicates
      const combined = [...prev, ...newSeats]
      return [...new Set(combined)]
    })
  }

  const clearGeneratedSeats = () => {
    setCurrentGeneratedSeats([])
    setAllGeneratedSeats([])
  }

  return {
    seats,
    loading,
    error,
    currentGeneratedSeats, // Latest generated seats (for showing in blue)
    allGeneratedSeats, // All generated seats (for showing assigned status)
    fetchSeats,
    addGeneratedSeats,
    clearGeneratedSeats,
  }
}