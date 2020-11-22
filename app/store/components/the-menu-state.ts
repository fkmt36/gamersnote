import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({
  name: 'components/the-menu-state',
  stateFactory: true,
  namespaced: true,
})
export default class TheMenuState extends VuexModule {
  private showMenu = false

  get getShowMenu(): boolean {
    return this.showMenu
  }

  @Mutation
  setShowMenu(showMenu: boolean): void {
    this.showMenu = showMenu
  }
}
