/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
  darkMode: ['class', '[data-theme="dim"]'],
  daisyui: {
    themes: ["corporate", "dim"],
  },
}
