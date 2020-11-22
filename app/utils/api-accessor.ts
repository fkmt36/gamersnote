import { ArticleApi, Configuration } from '@/api-client'

// eslint-disable-next-line import/no-mutable-exports
let $articleApi: () => ArticleApi

const clientConf = new Configuration({
  basePath: 'http://localhost:3000/api/v1',
})
const serverConf = new Configuration({
  basePath: 'http://localhost:3000/api/v1',
})
const articleCApi = new ArticleApi(clientConf)
const articleSApi = new ArticleApi(serverConf)

function initializeApi() {
  $articleApi = () => (process.client ? articleCApi : articleSApi)
}

export { initializeApi, $articleApi }
