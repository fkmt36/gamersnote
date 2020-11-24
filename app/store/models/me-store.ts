import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { User } from '@/api-client-axios'
import { $userApi } from '@/plugins/api'

@Module({
  name: 'models/me-store',
  stateFactory: true,
  namespaced: true,
})
export default class MeStore extends VuexModule {
  private me: User | null = null

  get getMe(): User | null {
    return this.me
  }

  @Mutation
  setMe(me: User | null) {
    this.me = me
  }

  @Action
  async fetchMe() {
    try {
      const res = await $userApi.getMe()
      this.setMe(res.data)
    } catch (e) {
      // if (isAxiosError(e) && e.response && e.response.status === 404) {
      //   isNewUserState.setIsNewUser(true)
      // }
      // TODO: エラースタックに積む
      this.setMe(null)
    }
  }

  @Action
  async postMe(user: User) {
    try {
      const res = await $userApi.postUser(user)
      this.setMe(res.data)
    } catch (e) {
      // TODO: エラースタックに積む
      this.setMe(null)
    }
  }
}
