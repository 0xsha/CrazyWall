export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,



  // Target: https://go.nuxtjs.dev/config-target
  target: "static",

  generate: {
    interval: 2000,
  },

  publicRuntimeConfig:{
    keys:{
      mapBox: process.env.MAPBOX_AUTH || "pk.eyJ1Ijoic2hhYWFhYSIsImEiOiJja3RyYjhvMzcxNGEyMm9vMjFzMXJmejZ0In0.vIbutACBG7oaqLD68rpDOg"

    }
  },



  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "CrazyWall",
    htmlAttrs: {
      lang: "en",
      class: ["dark"]
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      {
        rel: 'stylesheet',
        href: 'https://rsms.me/inter/inter.css'
      }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [], // '~/plugins/vue-fuse.js'

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    "@nuxtjs/tailwindcss",
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    "@nuxtjs/axios",
    "@nuxtjs/auth-next",
    "@nuxtjs/toast"
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    baseURL: process.env.BASE_URL ||  "https://api.crazywall.io"
  },

  router: {
    middleware: ["auth"]
  },

  auth: {

    strategies: {
      //cookie: {},
      local: {
        token: {
          property: "token",
          // required: true,
          // type: "Bearer"
        },
        user: {
           property: "user",
          // autoFetch: false
        },
        endpoints: {
          login: { url: "/api/v1/auth/login", method: "post" },
          logout: { url: "/api/v1/auth/logout", method: "post" },
          user: { url: "/api/v1/users/user", method: "get" }
        }
      }
    },
    // redirect: {
    //   // login: '/login',
    //   // logout: '/login',
    //   // callback: '/login',
    //   home: '/'
    // },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
};
