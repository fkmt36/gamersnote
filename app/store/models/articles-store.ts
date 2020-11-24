import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { Article } from '@/api-client-axios'
import { $articleApi } from '@/plugins/api'

@Module({
  name: 'models/articles-store',
  stateFactory: true,
  namespaced: true,
})
export default class ArticlesStore extends VuexModule {
  private articles: Array<Article> = []

  get getAllArticles(): Array<Article> {
    return this.articles
  }

  @Mutation
  setArticles(articles: Array<Article>) {
    this.articles.concat(articles)
  }

  @Action
  async fetchAllArticles() {
    try {
      const res = await $articleApi.getArticles()
      this.setArticles(res.data)
    } catch (err) {}
  }
}
