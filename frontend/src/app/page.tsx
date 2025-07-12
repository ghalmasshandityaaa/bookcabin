"use client"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Plane } from "lucide-react"
import { useCrewDetails } from "@/hooks/use-crew-details"
import { useNotifications } from "@/hooks/use-notifications"
import { useVoucherGeneration } from "@/hooks/use-voucher-generation"
import { useSeatManagement } from "@/hooks/use-seat-management"
import { CrewDetailsForm } from "@/components/forms/crew-details-form"
import { SeatMap } from "@/components/seat-map/seat-map"
import { PageHeader } from "@/components/layout/page-header"
import type { CrewDetails } from "@/types"

export default function VoucherAssignmentPage() {
  const { crewDetails, updateField } = useCrewDetails()
  const { error, success, setError, setSuccess, clearNotifications } = useNotifications()
  const { addGeneratedSeats, clearGeneratedSeats, currentGeneratedSeats, allGeneratedSeats } = useSeatManagement()

  const { generateVouchers, isLoading } = useVoucherGeneration({
    onSuccess: (message, seats) => {
      setSuccess(message)
      addGeneratedSeats(seats)
    },
    onError: setError,
  })

  const handleFieldChange = (field: keyof CrewDetails, value: string | Date | undefined) => {
    updateField(field, value)
    clearNotifications()

    // Clear generated seats when aircraft type changes
    if (field === "aircraft" && value !== crewDetails.aircraft) {
      clearGeneratedSeats()
    }
  }

  const handleGenerateVouchers = () => {
    generateVouchers(crewDetails)
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 p-4">
      <div className="max-w-6xl mx-auto space-y-6">
        <PageHeader />

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 lg:items-stretch">
          {/* Form Section */}
          <CrewDetailsForm
            crewDetails={crewDetails}
            onFieldChange={handleFieldChange}
            onSubmit={handleGenerateVouchers}
            isLoading={isLoading}
            error={error}
            success={success}
          />

          {/* Seat Map Section */}
          <Card className="shadow-lg">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Plane className="h-5 w-5 text-blue-600" />
                Seat Map
              </CardTitle>
              <CardDescription>Visual representation of seat assignments</CardDescription>
            </CardHeader>
            <CardContent>
              <SeatMap 
                aircraftType={crewDetails.aircraft} 
                currentGeneratedSeats={currentGeneratedSeats} 
                allGeneratedSeats={allGeneratedSeats} 
              />
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}