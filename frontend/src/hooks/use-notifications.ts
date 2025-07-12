"use client"

import { useState } from "react"
import type { NotificationState } from "@/types"

export const useNotifications = () => {
  const [notifications, setNotifications] = useState<NotificationState>({
    error: null,
    success: null,
  })

  const setError = (message: string) => {
    setNotifications({ error: message, success: null })
  }

  const setSuccess = (message: string) => {
    setNotifications({ error: null, success: message })
  }

  const clearNotifications = () => {
    setNotifications({ error: null, success: null })
  }

  return {
    ...notifications,
    setError,
    setSuccess,
    clearNotifications,
  }
}
