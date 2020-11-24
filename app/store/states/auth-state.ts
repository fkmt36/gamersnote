import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $fire } from '@/plugins/fire'
import { meStore } from '@/store'

@Module({
  name: 'states/auth-state',
  stateFactory: true,
  namespaced: true,
})
export default class AuthState extends VuexModule {
  private token: string = ''

  get getToken(): string {
    return this.token
  }

  @Mutation
  setToken(token: string): void {
    this.token = token
  }

  @Action
  async verifyEmailFromEmailLink(email: string, url: string) {
    try {
      const { user } = await $fire.auth.signInWithEmailLink(email, url)
      if (!user) throw new Error('user not found')
      const token = await user.getIdToken()
      this.setToken(token)
    } catch (e) {
      this.setToken('')
    }
  }

  @Action
  async onAuthStateChanged(params: any) {
    if (process.server) return
    try {
      if (!params || !params.authUser) throw new Error('user not found')
      const token = await params.authUser.getIdToken()
      this.setToken(token)
      await meStore.fetchMe()
    } catch (e) {
      this.setToken('')
      meStore.setMe(null)
    }
  }
}
