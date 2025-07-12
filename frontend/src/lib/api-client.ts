import type { CheckVoucherRequest, VoucherRequest, SeatsResponse, CheckResponse, GenerateResponse } from "@/types"

// API Configuration
const API_BASE_URL = "http://localhost:4000/api"

// HTTP Client utility
class ApiClient {
  private async request<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const url = `${API_BASE_URL}${endpoint}`

    const response = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
        ...options?.headers,
      },
      ...options,
    })

    if (!response.ok) {
      if (response.status === 500) {
        throw new Error("Internal server error")
      }

      const jsonResponse = await response.json()
      this.mapErrorResponse(jsonResponse)
    }

    const contentType = response.headers.get("content-type")
    if (!contentType || !contentType.includes("application/json")) {
      throw new Error("Server returned non-JSON response")
    }

    return response.json()
  }

  async getAircraftSeats(aircraftType: string): Promise<SeatsResponse> {
    return await this.request<SeatsResponse>(`/aircraft/seats?type=${encodeURIComponent(aircraftType)}`)
  }

  async checkVouchers(request: CheckVoucherRequest): Promise<CheckResponse> {
    return this.request<CheckResponse>("/check", {
      method: "POST",
      body: JSON.stringify(request),
    })
  }

  async generateVouchers(request: VoucherRequest): Promise<GenerateResponse> {
    return this.request<GenerateResponse>("/generate", {
      method: "POST",
      body: JSON.stringify(request),
    })
  }

  mapErrorResponse(jsonResponse: Record<string,string>) {
    if(jsonResponse.errors){
      if(jsonResponse.errors === 'seats/not-enough'){
        throw new Error(`Error: Seats is full in aircraft`)
      }
      throw new Error(`Error Request: ${jsonResponse.errors}`)
    }

    throw new Error(`HTTP error!`)
  }
}

export const apiClient = new ApiClient()
