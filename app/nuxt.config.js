export default {
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: 'app',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
    // reset.css
    '@/assets/destyle.css',
    // scss global variables
    // { src: '@/assets/global.scss', lang: 'scss' },
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [{ src: '@/plugins/api.ts' }],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    'nuxt-fontawesome',
    '@nuxtjs/firebase',
  ],

  firebase: {
    config: {
      apiKey: 'AIzaSyBMH9h_mtwYv3rgaDCLMSoheA8XHxS_VJA',
      authDomain: 'gamersnote-dev.firebaseapp.com',
      databaseURL: 'https://gamersnote-dev.firebaseio.com',
      projectId: 'gamersnote-dev',
      storageBucket: 'gamersnote-dev.appspot.com',
      messagingSenderId: '218913277344',
      appId: '1:218913277344:web:e84d70e30ecd392001d0c5',
      measurementId: 'G-K1CCGSY5CQ',
    },
    services: {
      auth: {
        persistence: 'local',
        ssr: false, // default
        // emulatorPort: 9099,
        // emulatorHost: 'http://localhost',
      },
    },
  },

  // Axios module configuration (https://go.nuxtjs.dev/config-axios)
  axios: {},

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {},

  fontawesome: {
    imports: [
      {
        set: '@fortawesome/free-solid-svg-icons',
        icons: ['fas'],
      },
      {
        set: '@fortawesome/free-brands-svg-icons',
        icons: ['fab'],
      },
    ],
  },
}
