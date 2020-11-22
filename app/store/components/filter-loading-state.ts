import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({
  name: 'components/filter-loading-state',
  stateFactory: true,
  namespaced: true,
})
export default class FilterLoadingState extends VuexModule {
  private isLoading = false
  private message = ''

  get getIsLoading(): boolean {
    return this.isLoading
  }

  get getMessage(): string {
    return this.message
  }

  @Mutation
  setLoading({
    isLoading,
    message,
  }: {
    isLoading: boolean
    message: string
  }): void {
    this.isLoading = isLoading
    this.message = message
  }

  @Mutation
  clearLoading() {
    this.isLoading = false
    this.message = ''
  }
}
