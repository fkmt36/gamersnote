import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({
  name: 'states/base-modal-state',
  stateFactory: true,
  namespaced: true,
})
export default class BaseModalState extends VuexModule {
  private showModal: boolean = false
  private message: string = ''
  private to: string = ''

  get getShowModal(): boolean {
    return this.showModal
  }

  get getMessage(): string {
    return this.message
  }

  get getTo(): string {
    return this.to
  }

  @Mutation
  setModal({
    showModal,
    message,
    to,
  }: {
    showModal: boolean
    message: string
    to?: string
  }): void {
    this.showModal = showModal
    this.message = message
    this.to = to || ''
  }
}
