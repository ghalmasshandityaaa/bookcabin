"use client"

import { useState } from "react"
import type { CrewDetails } from "@/types"

export const useCrewDetails = () => {
  const [crewDetails, setCrewDetails] = useState<CrewDetails>({
    name: "",
    id: "",
    flightNumber: "",
    date: undefined,
    aircraft: "",
  })

  const updateField = (field: keyof CrewDetails, value: string | Date | undefined) => {
    setCrewDetails((prev) => ({ ...prev, [field]: value }))
  }

  const resetForm = () => {
    setCrewDetails({
      name: "",
      id: "",
      flightNumber: "",
      date: undefined,
      aircraft: "",
    })
  }

  return {
    crewDetails,
    updateField,
    resetForm,
  }
}
