const colors = require('tailwindcss/colors')

module.exports = {
  purge: ['./public/*.html'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    colors: {
      // Build your palette here
      transparent: 'transparent',
      current: 'currentColor',
      gray: colors.warmGray,
      red: colors.red,
      blue: colors.sky,
      yellow: colors.amber,
    }
  },
  variants: {
    extend: {},
  },
  plugins: [],
}