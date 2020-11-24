import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import ArticlesStore from '@/store/models/articles-store'
import MeStore from '@/store/models/me-store'
import TheMenuState from '~/store/states/the-menu-state'
import FilterLoadingState from '~/store/states/filter-loading-state'
import FilterDarkenState from '~/store/states/filter-darken-state'
import AuthState from '~/store/states/auth-state'

// eslint-disable-next-line import/no-mutable-exports
let theMenuState: TheMenuState = new TheMenuState({})
// eslint-disable-next-line import/no-mutable-exports
let filterLoadingState: FilterLoadingState = new FilterLoadingState({})
// eslint-disable-next-line import/no-mutable-exports
let filterDarkenState: FilterDarkenState = new FilterDarkenState({})
// eslint-disable-next-line import/no-mutable-exports
let articlesStore: ArticlesStore = new ArticlesStore({})
// eslint-disable-next-line import/no-mutable-exports
let meStore: MeStore = new MeStore({})
// eslint-disable-next-line import/no-mutable-exports
let authState: AuthState = new AuthState({})

function initialiseStores(store: Store<any>): void {
  theMenuState = getModule(TheMenuState, store)
  filterLoadingState = getModule(FilterLoadingState, store)
  filterDarkenState = getModule(FilterDarkenState, store)
  articlesStore = getModule(ArticlesStore, store)
  meStore = getModule(MeStore, store)
  authState = getModule(AuthState, store)
}

export {
  initialiseStores,
  theMenuState,
  filterLoadingState,
  filterDarkenState,
  articlesStore,
  meStore,
  authState,
}
