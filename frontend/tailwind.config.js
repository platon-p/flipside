import colors from 'tailwindcss/colors';

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
        'primary': colors.orange[500]
      }
    },
  },
  plugins: [],
}

