import {
  Configuration,
  ArticleApi,
  UserApi,
  ImageApi,
  LikeApi,
  CommentApi,
} from '@/api-client-axios'

// eslint-disable-next-line import/no-mutable-exports
let $articleApi: () => ArticleApi
// eslint-disable-next-line import/no-mutable-exports
let $userApi: () => UserApi
// eslint-disable-next-line import/no-mutable-exports
let $imageApi: () => ImageApi
// eslint-disable-next-line import/no-mutable-exports
let $likeApi: () => LikeApi
// eslint-disable-next-line import/no-mutable-exports
let $commentApi: () => CommentApi

const initializeApi = () => {
  if (process.client) {
    const clientConf = new Configuration({
      basePath: '/api/v1',
    })
    $articleApi = () => new ArticleApi(clientConf)
    $userApi = () => new UserApi(clientConf)
    $imageApi = () => new ImageApi(clientConf)
    $likeApi = () => new LikeApi(clientConf)
    $commentApi = () => new CommentApi(clientConf)
  } else {
    const serverConf = new Configuration({
      basePath: process.env.API_URL + '/api/v1',
    })
    $articleApi = () => new ArticleApi(serverConf)
    $userApi = () => new UserApi(serverConf)
    $imageApi = () => new ImageApi(serverConf)
    $likeApi = () => new LikeApi(serverConf)
    $commentApi = () => new CommentApi(serverConf)
  }
}

initializeApi()

export { $articleApi, $userApi, $imageApi, $likeApi, $commentApi }
