/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.templ",
    "./templates/**/*_templ.go",
  ],
  theme: {
    extend: {
      colors: {
        navy: {
          DEFAULT: "#1e3a5f",
          light: "#2b4f82",
          dark: "#0f2440",
          50: "#eaf1f8",
          100: "#c8d9ed",
          200: "#92b3da",
          300: "#5c8dc7",
          400: "#3670b0",
          500: "#1e3a5f",
          600: "#1a3355",
          700: "#152a47",
          800: "#0f2440",
          900: "#081530",
        },
        sage: {
          DEFAULT: "#3a8a5c",
          light: "#4dab74",
          dark: "#2d6b47",
          50: "#edf7f1",
          100: "#cdebd8",
          200: "#9bd7b1",
          300: "#69c38a",
          400: "#4dab74",
          500: "#3a8a5c",
          600: "#2d6b47",
          700: "#215132",
          800: "#163822",
          900: "#0b1f13",
        },
        card: {
          orange: "#e8623f",
          yellow: "#f2c94c",
          blue: "#4a90d9",
          dark: "#1e2a38",
          teal: "#2bbaa0",
          purple: "#6c5ce7",
        },
        surface: {
          DEFAULT: "#f5f3f0",
          raised: "#faf9f7",
        },
        charcoal: {
          DEFAULT: "#1a1a2e",
          light: "#6b7280",
        },
      },
      fontFamily: {
        serif: ['"Libre Baskerville"', "Georgia", "serif"],
        sans: ['"Inter"', "system-ui", "sans-serif"],
      },
      lineHeight: {
        relaxed: "1.75",
      },
      borderRadius: {
        "2xl": "1rem",
        "3xl": "1.5rem",
        "4xl": "2rem",
      },
      boxShadow: {
        "glass": "0 8px 32px rgba(0,0,0,0.06)",
        "glass-lg": "0 12px 48px rgba(0,0,0,0.08)",
        "card-hover": "0 20px 40px rgba(0,0,0,0.12)",
        "inner-glow": "inset 0 1px 0 rgba(255,255,255,0.1)",
        "soft": "0 2px 8px rgba(0,0,0,0.04), 0 1px 2px rgba(0,0,0,0.06)",
        "soft-lg": "0 4px 16px rgba(0,0,0,0.06), 0 2px 4px rgba(0,0,0,0.04)",
      },
      animation: {
        "fade-up": "fadeUp 0.5s ease-out both",
        "fade-up-delay": "fadeUp 0.5s ease-out 0.1s both",
        "scale-in": "scaleIn 0.3s ease-out both",
      },
      keyframes: {
        fadeUp: {
          "0%": { opacity: "0", transform: "translateY(12px)" },
          "100%": { opacity: "1", transform: "translateY(0)" },
        },
        scaleIn: {
          "0%": { opacity: "0", transform: "scale(0.95)" },
          "100%": { opacity: "1", transform: "scale(1)" },
        },
      },
      typography: ({ theme }) => ({
        DEFAULT: {
          css: {
            "--tw-prose-body": theme("colors.charcoal.DEFAULT"),
            "--tw-prose-headings": theme("colors.navy.dark"),
            "--tw-prose-links": theme("colors.sage.DEFAULT"),
            "--tw-prose-bold": theme("colors.charcoal.DEFAULT"),
            "--tw-prose-code": theme("colors.navy.DEFAULT"),
            lineHeight: "1.8",
            h1: { fontFamily: theme("fontFamily.serif").join(", "), letterSpacing: "-0.02em" },
            h2: { fontFamily: theme("fontFamily.serif").join(", "), letterSpacing: "-0.01em" },
            h3: { fontFamily: theme("fontFamily.serif").join(", ") },
            h4: { fontFamily: theme("fontFamily.serif").join(", ") },
            a: {
              color: theme("colors.sage.DEFAULT"),
              textDecorationThickness: "1px",
              textUnderlineOffset: "3px",
              transition: "color 0.2s ease",
              "&:hover": { color: theme("colors.sage.light") },
            },
            code: {
              backgroundColor: "#f0eeeb",
              padding: "0.2rem 0.4rem",
              borderRadius: "0.375rem",
              fontWeight: "400",
              fontSize: "0.875em",
            },
            "code::before": { content: '""' },
            "code::after": { content: '""' },
            img: { borderRadius: "1rem" },
            blockquote: {
              borderLeftColor: theme("colors.sage.DEFAULT"),
              borderLeftWidth: "3px",
              fontStyle: "italic",
              backgroundColor: theme("colors.sage.50"),
              padding: "1rem 1.5rem",
              borderRadius: "0 0.75rem 0.75rem 0",
            },
            hr: {
              borderColor: "#e5e5e5",
            },
            "pre": {
              borderRadius: "1rem",
            },
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
