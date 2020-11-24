import { Configuration, ArticleApi, UserApi } from '@/api-client-axios'
import { authState } from '@/store'

// eslint-disable-next-line import/no-mutable-exports
let $articleApi: ArticleApi
// eslint-disable-next-line import/no-mutable-exports
let $userApi: UserApi

const initializeApi = () => {
  if (process.client) {
    const clientConf = new Configuration({
      basePath: 'http://localhost:3000/api/v1',
      apiKey: () => authState.getToken,
    })
    $articleApi = new ArticleApi(clientConf)
    $userApi = new UserApi(clientConf)
  } else {
    const serverConf = new Configuration({
      basePath: 'http://nginx/api/v1',
      apiKey: () => authState.getToken,
    })
    $articleApi = new ArticleApi(serverConf)
    $userApi = new UserApi(serverConf)
  }
}

initializeApi()

export { $articleApi, $userApi }
