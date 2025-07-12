import type { CrewDetails } from "@/types"

export interface ValidationError {
  field: keyof CrewDetails
  message: string
}

export const validateCrewDetails = (details: CrewDetails): string | null => {
  // Name validation - minimum 2 characters
  if (!details.name.trim()) {
    return "Crew name is required"
  }
  if (details.name.trim().length < 2) {
    return "Crew name must be at least 2 characters"
  }

  // ID validation - must be numeric string
  if (!details.id.trim()) {
    return "Crew ID is required"
  }
  if (!/^\d+$/.test(details.id.trim())) {
    return "Crew ID must contain only numbers"
  }

  // Flight number validation - minimum 3 characters
  if (!details.flightNumber.trim()) {
    return "Flight number is required"
  }
  if (details.flightNumber.trim().length < 3) {
    return "Flight number must be at least 3 characters"
  }

  // Date validation
  if (!details.date) {
    return "Flight date is required"
  }

  // Aircraft validation
  if (!details.aircraft) {
    return "Aircraft type is required"
  }

  return null
}

export const isValidDateFormat = (date: Date): boolean => {
  // Check if date is valid and not in the past
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  const selectedDate = new Date(date)
  selectedDate.setHours(0, 0, 0, 0)

  return selectedDate >= today
}
