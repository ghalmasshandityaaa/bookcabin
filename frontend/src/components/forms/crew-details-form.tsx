"use client"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Loader2, CheckCircle } from "lucide-react"
import { type CrewDetails, AIRCRAFT_TYPES } from "@/types"
import { DatePicker } from "../ui/date-picker"
import { NotificationAlerts } from "../ui/notification-alerts"

interface CrewDetailsFormProps {
  crewDetails: CrewDetails
  onFieldChange: (field: keyof CrewDetails, value: string | Date | undefined) => void
  onSubmit: () => void
  isLoading: boolean
  error: string | null
  success: string | null
}

export function CrewDetailsForm({
  crewDetails,
  onFieldChange,
  onSubmit,
  isLoading,
  error,
  success,
}: CrewDetailsFormProps) {
  return (
    <Card className="shadow-lg flex flex-col">
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <CheckCircle className="h-5 w-5 text-green-600" />
          Flight & Crew Details
        </CardTitle>
        <CardDescription>Enter the required information to generate voucher seat assignments</CardDescription>
      </CardHeader>
      <CardContent className="space-y-4 flex flex-col flex-1">
        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="crewName">Crew Name</Label>
            <Input
              id="crewName"
              placeholder="Enter crew name"
              value={crewDetails.name}
              onChange={(e) => onFieldChange("name", e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="crewId">Crew ID</Label>
            <Input
              id="crewId"
              placeholder="Enter crew ID"
              value={crewDetails.id}
              onChange={(e) => onFieldChange("id", e.target.value)}
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="flightNumber">Flight Number</Label>
            <Input
              id="flightNumber"
              placeholder="e.g., GA102"
              value={crewDetails.flightNumber}
              onChange={(e) => onFieldChange("flightNumber", e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <Label>Flight Date</Label>
            <DatePicker
              date={crewDetails.date}
              onSelect={(date) => onFieldChange("date", date)}
              placeholder="Pick a date"
            />
          </div>
        </div>

        <div className="space-y-2">
          <Label htmlFor="aircraft">Aircraft Type</Label>
          <Select value={crewDetails.aircraft} onValueChange={(value) => onFieldChange("aircraft", value)}>
            <SelectTrigger>
              <SelectValue placeholder="Select aircraft type" />
            </SelectTrigger>
            <SelectContent>
              {AIRCRAFT_TYPES.map((aircraft) => (
                <SelectItem key={aircraft.value} value={aircraft.value}>
                  {aircraft.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        {/* Spacer to push content to bottom */}
        <div className="flex-1 flex flex-col justify-end space-y-4">
          <NotificationAlerts error={error} success={success} />

          {/* Generate Button - At Bottom */}
          <Button onClick={onSubmit} disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700" size="lg">
            {isLoading ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Generating Vouchers...
              </>
            ) : (
              "Generate Vouchers"
            )}
          </Button>
        </div>
      </CardContent>
    </Card>
  )
}
