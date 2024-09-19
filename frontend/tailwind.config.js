/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#d8dbe2",
        secondary: "#a9bcd0",
        accent: "#58a4b0",
        accentDark: "#407e87",
        txtColor: "#373f51",
        brdr: "#1b1b1e",
      },
    },
  },
  plugins: [],
};
