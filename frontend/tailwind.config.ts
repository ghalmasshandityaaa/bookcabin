import type { Config } from 'tailwindcss'

const config: Config = {
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
      },
    },
  },
  content: ['./src/**/*.{js,ts,jsx,tsx}'],
  plugins: [],
}

export default config
