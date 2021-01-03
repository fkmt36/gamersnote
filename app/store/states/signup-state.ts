import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({
  name: 'states/signup-state',
  stateFactory: true,
  namespaced: true,
})
export default class SignupState extends VuexModule {
  private email: string = ''

  get getEmail(): string {
    return this.email
  }

  @Mutation
  setEmail(email: string): void {
    this.email = email
  }
}
