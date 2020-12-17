import {
  Configuration,
  ArticleApi,
  UserApi,
  ImageApi,
} from '@/api-client-axios'

console.log('api', process.client)

// eslint-disable-next-line import/no-mutable-exports
let $articleApi: () => ArticleApi
// eslint-disable-next-line import/no-mutable-exports
let $userApi: () => UserApi
// eslint-disable-next-line import/no-mutable-exports
let $imageApi: () => ImageApi

const initializeApi = () => {
  if (process.client) {
    const clientConf = new Configuration({
      basePath: process.env.API_URL_BROWSER,
    })
    $articleApi = () => new ArticleApi(clientConf)
    $userApi = () => new UserApi(clientConf)
    $imageApi = () => new ImageApi(clientConf)
  } else {
    const serverConf = new Configuration({
      basePath: process.env.API_URL,
    })
    $articleApi = () => new ArticleApi(serverConf)
    $userApi = () => new UserApi(serverConf)
    $imageApi = () => new ImageApi(serverConf)
  }
}

initializeApi()

export { $articleApi, $userApi, $imageApi }
