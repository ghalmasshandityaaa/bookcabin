# 🌐 BookCabin Frontend - Airline Voucher Seat Assignment

> Modern React-based frontend for airline crew to generate voucher seat assignments with intuitive UI and real-time validation.

## 📋 Overview

The BookCabin frontend is a Next.js application that provides airline crew members with an easy-to-use interface for generating random seat vouchers. The app features form validation, aircraft-specific configurations, and seamless integration with the backend API.

## ✨ Features

- **Modern UI/UX** - Clean, responsive design optimized for airline crew workflows
- **Real-time Validation** - Instant feedback on form inputs and data validation
- **Aircraft Selection** - Dropdown with ATR, Airbus 320, and Boeing 737 Max options
- **Duplicate Prevention** - Checks for existing vouchers before generation
- **Responsive Design** - Works seamlessly on desktop, tablet, and mobile devices
- **Error Handling** - User-friendly error messages and loading states
- **TypeScript Support** - Type safety and better development experience

## 🛠️ Tech Stack

- **Framework**: Next.js 15+ (React 18+)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: React Hooks (useState, useEffect)
- **HTTP Client**: Fetch API
- **Font**: Geist (Optimized by Next.js)
- **Package Manager**: Yarn

## 🚀 Getting Started

### Prerequisites
- **Node.js** (v18 or higher)
- **Yarn** (v1.22 or higher)
- **Backend API** running on `http://localhost:3000`

### Installation

1. **Navigate to frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   yarn install
   ```

3. **Start development server**
   ```bash
   yarn dev
   ```

4. **Access the application**
   Open [http://localhost:3000](http://localhost:3000) in your browser

## 📁 Project Structure

```
frontend/
├── .next/                  # Next.js build output (auto-generated)
├── node_modules/           # Node.js dependencies (auto-generated)
├── src/                    # Source code directory
│   ├── app/               # Next.js App Router
│   │   ├── globals.css    # Global styles
│   │   ├── layout.tsx     # Root layout component
│   │   └── page.tsx       # Home page component
│   ├── components/        # Reusable UI components
│   │   ├── forms/        # Form-related components
│   │   ├── ui/           # Generic UI components
│   │   └── layout/       # Layout components
│   ├── lib/              # Utility functions
│   │   ├── api.ts        # API client functions
│   │   ├── utils.ts      # Helper utilities
│   │   └── constants.ts  # App constants
│   └── types/            # TypeScript type definitions
│       └── index.ts      # Shared types
├── public/                # Static assets
├── .gitignore            # Git ignore rules
├── components.json       # shadcn/ui components config
├── eslint.config.mjs     # ESLint configuration
├── next.config.ts        # Next.js configuration
├── next-env.d.ts         # Next.js TypeScript definitions
├── package.json          # Dependencies and scripts
├── postcss.config.mjs    # PostCSS configuration
├── tailwind.config.ts    # Tailwind CSS configuration
├── tsconfig.json         # TypeScript configuration
└── README.md             # This file
```

## 🎯 Core Components

### Main Form Component
The primary interface where crew members input:
- **Crew Name** - Full name of the crew member
- **Crew ID** - Unique identifier for the crew member
- **Flight Number** - Aircraft flight identifier
- **Flight Date** - Date in DD-MM-YY format
- **Aircraft Type** - Dropdown selection (ATR, Airbus 320, Boeing 737 Max)

### Validation System
- **Client-side validation** for all form fields
- **Date format validation** (DD-MM-YY)
- **Flight number format validation**
- **Required field validation**

### Results Display
- **Success state** - Shows the 3 generated seat numbers
- **Error state** - Displays user-friendly error messages
- **Loading state** - Indicates API request in progress
