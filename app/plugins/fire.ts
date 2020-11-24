import { Plugin } from '@nuxt/types'
import { NuxtFireInstance } from '@nuxtjs/firebase'

// eslint-disable-next-line import/no-mutable-exports
let $fire: NuxtFireInstance

const initializeFire: Plugin = ({ $fire: fire }) => {
  $fire = fire
}

export { $fire }
export default initializeFire
