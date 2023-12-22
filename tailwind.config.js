/** @type {import('tailwindcss').Config} */
export const content = ["internal/views/**/*.templ"];
export const darkMode = "class";
export const theme = {
  extend: {
    fontFamily: {
      mono: ["Courier Prime", "monospace"],
    },
  },
};
export const plugins = [];
export const corePlugins = {
  preflight: true,
};
