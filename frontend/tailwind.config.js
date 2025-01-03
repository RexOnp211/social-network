/** @type {import('tailwindcss').Config} */
const plugin = require("tailwindcss/plugin");

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
        error: "ef4444",
        valid: "22c55e",
      },
    },
  },
  plugins: [
    plugin(function ({ addUtilities }) {
      addUtilities({
        ".attentionBG": {
          "background-color": "#ffedd5",
          padding: "1rem",
          "border-radius": "0.5rem",
          "margin-top": "1rem",
          width: "100%",
        },
        ".basicButton": {
          "margin-top": "0.5rem", // mt-2
          transition: "all 0.3s ease-in", // transition-colors ease-in
          "background-color": "#58a4b0", // bg-accent
          color: "#ffffff", // text-white
          "border-radius": "0.5rem", // rounded-lg
          padding: "0.5rem", // p-2
          width: "100%", // w-full
        },
        ".customButton:hover": {
          "background-color": "#407e87", // hover:bg-accentDark
        },
      });
    }),
  ],
};
