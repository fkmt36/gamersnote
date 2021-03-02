export default {
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: 'GamersNote',
    prefix:
      'og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# website: http://ogp.me/ns/website#',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { property: 'og:url', content: process.env.APP_URL },
      { property: 'og:type', content: 'website' },
      { property: 'og:title', content: 'GamersNote' },
      {
        property: 'og:description',
        content:
          'GamersNoteはゲームに関する記事を投稿して交流できるメディアプラットフォームです。',
      },
      { property: 'og:site_name', content: 'GamersNote' },
      { property: 'og:image', content: process.env.S3_URL + '/ogp.png' },
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
  modules: ['nuxt-fontawesome'],

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
