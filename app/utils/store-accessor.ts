import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import TheMenuState from '@/store/components/the-menu-state'
import FilterLoadingState from '@/store/components/filter-loading-state'
import FilterDarkenState from '@/store/components/filter-darken-state'
import ArticlesStore from '@/store/models/articles-store'

// eslint-disable-next-line import/no-mutable-exports
let theMenuState: TheMenuState = new TheMenuState({})
// eslint-disable-next-line import/no-mutable-exports
let filterLoadingState: FilterLoadingState = new FilterLoadingState({})
// eslint-disable-next-line import/no-mutable-exports
let filterDarkenState: FilterDarkenState = new FilterDarkenState({})
// eslint-disable-next-line import/no-mutable-exports
let articlesStore: ArticlesStore = new ArticlesStore({})

function initialiseStores(store: Store<any>): void {
  theMenuState = getModule(TheMenuState, store)
  filterLoadingState = getModule(FilterLoadingState, store)
  filterDarkenState = getModule(FilterDarkenState, store)
  articlesStore = getModule(ArticlesStore, store)
}

export {
  initialiseStores,
  theMenuState,
  filterLoadingState,
  filterDarkenState,
  articlesStore,
}
