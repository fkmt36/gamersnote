import { ActionContext } from 'vuex/types'
import { Context } from '@nuxt/types'
import { meStore, authState } from '@/store'

export async function nuxtServerInit(
  _: ActionContext<any, any>,
  { res }: Context
) {
  if (res && res.locals && res.locals.user) {
    authState.setToken(res.locals.user.idToken)
    await meStore.fetchMe()
  }
}
