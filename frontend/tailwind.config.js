/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    clipPath: {
      mypolygon: "polygon( 0 0 ,100% 0,100% 80%,0 100%)",
    },
    extend: {
      colors: {
        primary: "#202124",
        secondary: "#1B1C20",
        gradientFrom: "#bd2a96",
        gradientVia: "#EA2940",
        gradientTo: "#DAA827",
      }
    },
  },
  plugins: [
    require('tailwind-clip-path'),
  ],
}

