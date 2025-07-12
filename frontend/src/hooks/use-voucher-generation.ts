"use client"

import { useState } from "react"
import { apiClient } from "@/lib/api-client"
import type { CrewDetails, VoucherRequest } from "@/types"
import { formatDateForAPI } from "@/lib/utils/date"
import { validateCrewDetails } from "@/lib/utils/validation"

interface UseVoucherGenerationProps {
  onSuccess: (message: string, seats: string[]) => void
  onError: (message: string) => void
  // Remove onClearSeats since we don't want to clear all seats
}

export const useVoucherGeneration = ({ onSuccess, onError }: UseVoucherGenerationProps) => {
  const [isLoading, setIsLoading] = useState(false)

  const generateVouchers = async (crewDetails: CrewDetails) => {
    // Validate form
    const validationError = validateCrewDetails(crewDetails)
    if (validationError) {
      onError(validationError)
      return
    }

    setIsLoading(true)

    try {
      // Step 1: Check if vouchers already exist
      const checkResult = await apiClient.checkVouchers({
        flightNumber: crewDetails.flightNumber,
        date: formatDateForAPI(crewDetails.date!),
      })

      if (checkResult.success && checkResult.exists) {
        onError("Vouchers have already been generated for this flight and date")
        return
      }

      // Step 2: Generate new vouchers (don't clear existing seats)
      const generateRequest: VoucherRequest = {
        name: crewDetails.name,
        id: crewDetails.id,
        flightNumber: crewDetails.flightNumber,
        date: formatDateForAPI(crewDetails.date!),
        aircraft: crewDetails.aircraft,
      }

      const generateResult = await apiClient.generateVouchers(generateRequest)

      if (!generateResult.success) {
        if (generateResult.errors === "vouchers/already-exists") {
          onError("Vouchers have already been generated for this flight and date")
        } else {
          onError("Failed to generate vouchers")
        }
        return
      }

      // Step 3: Success
      if (generateResult.seats && Array.isArray(generateResult.seats)) {
        const seats = generateResult.seats
        onSuccess(`Successfully generated ${seats.length} voucher seats: ${seats.join(", ")}`, seats)
      }
    } catch (error) {
      if (error instanceof Error && error.message === "Internal server error") {
        onError("Internal server error")
      } else {
        onError(error instanceof Error ? error.message : "An unexpected error occurred")
      }
    } finally {
      setIsLoading(false)
    }
  }

  return {
    generateVouchers,
    isLoading,
  }
}