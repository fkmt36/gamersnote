import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({
  name: 'states/filter-darken-state',
  stateFactory: true,
  namespaced: true,
})
export default class FilterDarkenState extends VuexModule {
  private isDarken: boolean = false

  get getIsDarken(): boolean {
    return this.isDarken
  }

  @Mutation
  setDarken(isDarken: boolean): void {
    this.isDarken = isDarken
  }
}
