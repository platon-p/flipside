/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': 'inter, Arial'
      },
      colors: {
        'primary': '#EA580B'
      }
    },
  },
  plugins: [],
}

