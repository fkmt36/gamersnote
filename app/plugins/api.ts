import {
  Configuration,
  ArticleApi,
  UserApi,
  ImageApi,
} from '@/api-client-axios'

// eslint-disable-next-line import/no-mutable-exports
let $articleApi: () => ArticleApi
// eslint-disable-next-line import/no-mutable-exports
let $userApi: () => UserApi
// eslint-disable-next-line import/no-mutable-exports
let $imageApi: () => ImageApi

const initializeApi = () => {
  if (process.client) {
    const clientConf = new Configuration({
      basePath: '/api/v1',
    })
    console.log('BaseURL', clientConf.basePath)
    $articleApi = () => new ArticleApi(clientConf)
    $userApi = () => new UserApi(clientConf)
    $imageApi = () => new ImageApi(clientConf)
  } else {
    const serverConf = new Configuration({
      basePath: process.env.API_URL + '/api/v1',
    })
    console.log('BaseURL', serverConf.basePath)
    $articleApi = () => new ArticleApi(serverConf)
    $userApi = () => new UserApi(serverConf)
    $imageApi = () => new ImageApi(serverConf)
  }
}

initializeApi()

export { $articleApi, $userApi, $imageApi }
