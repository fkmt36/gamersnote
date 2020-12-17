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
  private currentArticle: Article | null = null

  get getAllArticles(): Array<Article> {
    return this.articles
  }

  get getArticleByID(): (articleID: string) => Article | null {
    return (articleID: string) => {
      return this.articles.find((e) => e.article_id === articleID) || null
    }
  }

  get getCurrentArticle(): Article | null {
    return this.currentArticle
  }

  @Mutation
  setArticles(articles: Array<Article>) {
    this.articles = articles
  }

  @Mutation
  concatArticles(articles: Array<Article>) {
    this.articles = this.articles.concat(articles)
  }

  @Mutation
  setCurrentArticle(article: Article | null) {
    this.currentArticle = article
  }

  // @Action
  // async fetchCurrentArticle(articleID: string) {
  //   try {
  //     this.setCurrentArticle(this.getArticleByID(articleID))
  //     if (this.currentArticle) return
  //     this.setCurrentArticle((await $articleApi.getArticle(articleID)).data)
  //   } catch {}
  // }

  // @Action
  // async fetchAllArticles() {
  //   try {
  //     const res = await $articleApi.getArticles()
  //     this.setArticles(res.data)
  //   } catch (err) {}
  // }

  // @Action
  // async postArticle(article: Article): Promise<Article | null> {
  //   try {
  //     const res = await $articleApi.postArticle(article)
  //     return res.data
  //   } catch {
  //     // エラースタック
  //     return null
  //   }
  // }
}
