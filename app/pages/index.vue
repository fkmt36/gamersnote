<template>
  <div class="home">
    <div class="action" @click="openModal = !openModal">
      <span class="sort-title">{{ order }}</span>
      <font-awesome-icon :icon="['fas', 'sort']" />
    </div>
    <!-- <div v-show="openModal" class="modal">
      <div><button @click="initLatest">最新</button></div>
      <div><button @click="initPopular">人気</button></div>
    </div> -->
    <div v-show="openModal" class="modal">
      <div><NuxtLink to="/">最新</NuxtLink></div>
      <div>
        <NuxtLink to="/?sort=popular">人気</NuxtLink>
      </div>
    </div>
    <div class="article-list">
      <ListArticle :articles="articles" />
    </div>
    <client-only>
      <infinite-loading @infinite="infiniteHandler"></infinite-loading>
    </client-only>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import ListArticle from '@/components/ListArticle.vue'
import InfiniteLoading from 'vue-infinite-loading'
import { $articleApi } from '@/plugins/api'
import { Article } from '~/api-client-axios'

interface Data {
  articles: Array<Article>
  order: string
  openModal: boolean
  offset: number
}

export default Vue.extend({
  key(route) {
    return route.fullPath
  },
  components: {
    ListArticle,
    InfiniteLoading,
  },
  async asyncData({ query }): Promise<Data> {
    try {
      if (query.sort === 'popular') {
        const { data } = await $articleApi().getArticles(-1, '', 'popular')
        return {
          articles: data,
          order: '人気',
          openModal: false,
          offset: -1,
        }
      } else {
        const { data } = await $articleApi().getArticles(-1, '', '')
        return {
          articles: data,
          order: '最新',
          openModal: false,
          offset: -1,
        }
      }
    } catch (e) {
      return {
        articles: [],
        order: '最新',
        openModal: false,
        offset: -1,
      }
    }
  },
  data() {
    return {} as Data
  },
  async mounted() {
    try {
      if (this.$route.query.sort === 'popular' && this.order === '最新') {
        this.order = '人気'

        this.offset = -1
        this.openModal = false
        const { data } = await $articleApi().getArticles(-1, '', 'popular')
        this.articles = data
      } else if (
        this.$route.query.sort !== 'popular' &&
        this.order === '人気'
      ) {
        this.order = '最新'

        this.offset = -1
        this.openModal = false
        const { data } = await $articleApi().getArticles(-1, '', '')
        this.articles = data
      }
    } catch {
      this.articles = []
    }
  },
  methods: {
    async infiniteHandler($state: any) {
      if (this.order === '人気') {
        this.offset += 20
        const { data } = await $articleApi().getArticles(
          this.offset,
          '',
          'popular'
        )
        if (data.length) {
          this.articles = this.articles.concat(data)
          $state.loaded()
        } else {
          $state.complete()
        }
      } else {
        this.offset += 20
        const { data } = await $articleApi().getArticles(this.offset, '', '')
        if (data.length) {
          this.articles = this.articles.concat(data)
          $state.loaded()
        } else {
          $state.complete()
        }
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.home {
  max-width: 1440px;
  margin: 0 auto;
  position: relative;
}

.modal {
  width: 180px;
  height: 80px;
  background-color: white;
  position: absolute;
  right: 0;
  border-radius: 5px;
  border: 1px solid $stroke-grey;
  text-align: center;
  z-index: 2;
  padding: 10px;
  div {
    margin: auto 0;
    font-size: 1rem;
  }
}

.action {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  height: 40px;
  font-weight: bold;
  color: $dark-grey;
  padding: 0 15px;

  .sort-title {
    margin-right: 5px;
  }
}
</style>
