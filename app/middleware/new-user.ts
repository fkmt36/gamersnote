import { Context } from '@nuxt/types'
// なぜかimportできない
// import { meStore, authState } from '@/store'

export default function ({ redirect, store, route }: Context) {
  const isNewUser =
    store.getters['states/auth-state/getToken'] &&
    !store.getters['models/me-store/getMe']
  if (route.path !== '/welcome' && isNewUser) {
    redirect('/welcome')
  }
  if (route.path === '/welcome' && !isNewUser) {
    redirect('/')
  }
}
