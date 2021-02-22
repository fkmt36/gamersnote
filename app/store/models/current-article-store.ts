import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { Article } from '@/api-client-axios'
import { $articleApi, $likeApi } from '@/plugins/api'

@Module({
  name: 'models/current-article-store',
  stateFactory: true,
  namespaced: true,
})
export default class CurrentArticleStore extends VuexModule {
  private articleId: string = ''
  private thumbnail: string = ''
  private title: string = ''
  private body: string = ''
  private avatar: string = ''
  private username: string = ''
  private createdAt: string = ''
  private likeCount: number = 0
  private liked: boolean = false

  get getArticleId(): string {
    return this.articleId
  }

  get getThumbnail(): string {
    return this.thumbnail
  }

  get getTitle(): string {
    return this.title
  }

  get getBody(): string {
    return this.body
  }

  get getAvatar(): string {
    return this.avatar
  }

  get getUsername(): string {
    return this.username
  }

  get getCreatedAt(): string {
    return this.createdAt
  }

  get getLikeCount(): number {
    return this.likeCount
  }

  get getLiked(): boolean {
    return this.liked
  }

  @Mutation
  setArticle(article: Article) {
    this.articleId = article.article_id || ''
    this.thumbnail = article.thumbnail_url
    this.title = article.title
    this.body = article.body
    this.avatar = article.author?.avatar_url || ''
    this.username = article.author?.username || ''
    this.createdAt = article.created_at || ''
    this.likeCount = article.like_count || 0
  }

  @Mutation
  incrementLikeCount() {
    this.likeCount += 1
  }

  @Mutation
  decrementLikeCount() {
    this.likeCount -= 1
  }

  @Mutation
  setLiked(liked: boolean) {
    this.liked = liked
  }

  @Action({ rawError: true })
  async like() {
    await $likeApi().putLike(this.articleId)
    this.incrementLikeCount()
    this.setLiked(true)
  }

  @Action({ rawError: true })
  async dislike() {
    await $likeApi().deleteLike(this.articleId)
    this.decrementLikeCount()
    this.setLiked(false)
  }

  @Action({ rawError: true })
  async fetchArticle(articleId: string, cookie: string | undefined) {
    if (process.server) {
      const { data } = await $articleApi().getArticle(articleId)
      const liked = await (async (): Promise<boolean> => {
        try {
          await $likeApi().getLike(articleId, {
            headers: {
              Cookie: cookie,
            },
          })
          return true
        } catch {
          return false
        }
      })()
      this.setArticle(data)
      this.setLiked(liked)
    } else {
      const { data } = await $articleApi().getArticle(articleId)
      const liked = await (async (): Promise<boolean> => {
        try {
          await $likeApi().getLike(articleId)
          return true
        } catch {
          return false
        }
      })()
      this.setArticle(data)
      this.setLiked(liked)
    }
  }

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
