import { Store } from 'vuex'
import { initialiseStores } from '@/utils/store-accessor'
import { nuxtServerInit } from '@/utils/nuxt-server-init'

const initializer = (store: Store<any>) => initialiseStores(store)
export const actions = { nuxtServerInit }
export const plugins = [initializer]
export * from '@/utils/store-accessor'
