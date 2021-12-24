
///const defaultTheme = require("tailwindcss/defaultTheme");


module.exports = {
  purge: [],
  darkMode: 'class', // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        draculaBackground: '#282a36',
        draculaForeground: '#f8f8f2',
        draculaComment: '#6272a4',
        draculaSelection: '#44475a',
        draculaRed: '#ff5555',
        draculaYellow: '#f1fa8c',
        draculaGreen: '#50fa7b',
        draculaOrange: '#ffb86c',
        draculaPink: '#ff79c6',
        draculaPurple: '#bd93f9',
        glowPurple: '#E2CAFA'
      }
    }
  },
  variants: {
    extend: {

      opacity: ['disabled'],

      // fontFamily: {
      //   sans: ['Inter var', ...defaultTheme.fontFamily.sans],
      // },
    },
  },
  plugins: [],
}
