// Core domain types
export interface Seat {
  row_number: number
  seat: string
  assigned: boolean
  generated?: boolean
}

export interface CrewDetails {
  name: string
  id: string
  flightNumber: string
  date: Date | undefined
  aircraft: string
}

export interface VoucherRequest {
  name: string
  id: string
  flightNumber: string
  date: string
  aircraft: string
}

export interface CheckVoucherRequest {
  flightNumber: string
  date: string
}

// API Response types
export interface ApiResponse<T = any> {
  success: boolean
  data?: T
  error?: string
  errors?: string
}

export interface SeatsResponse extends ApiResponse {
  data: Seat[]
}

export interface CheckResponse extends ApiResponse {
  exists: boolean
}

export interface GenerateResponse extends ApiResponse {
  seats: string[]
}

// UI State types
export interface NotificationState {
  error: string | null
  success: string | null
}

export interface LoadingState {
  isLoading: boolean
}

// Constants
export const AIRCRAFT_TYPES = [
  { value: "ATR", label: "ATR" },
  { value: "Airbus 320", label: "Airbus 320" },
  { value: "Boeing 737 Max", label: "Boeing 737 Max" },
] as const

export type AircraftType = (typeof AIRCRAFT_TYPES)[number]["value"]
