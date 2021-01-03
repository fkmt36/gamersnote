import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({
  name: 'states/base-modal-state',
  stateFactory: true,
  namespaced: true,
})
export default class BaseModalState extends VuexModule {
  private showModal: boolean = false
  private message: string = ''

  get getShowModal(): boolean {
    return this.showModal
  }

  get getMessage(): string {
    return this.message
  }

  @Mutation
  setModal({
    showModal,
    message,
  }: {
    showModal: boolean
    message: string
  }): void {
    this.showModal = showModal
    this.message = message
  }
}
