<template>
  <div class="home">
    <div class="action">
      <span class="sort-title">最新</span>
      <font-awesome-icon :icon="['fas', 'sort']" />
    </div>
    <div class="article-list">
      <ListArticle :articles="articles" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import ListArticle from '@/components/ListArticle.vue'
import { $articleApi } from '@/plugins/api'
import { Article } from '~/api-client-axios'

interface Data {
  articles: Array<Article>
}

export default Vue.extend({
  components: {
    ListArticle,
  },
  async asyncData(): Promise<Data> {
    try {
      const result = await $articleApi().getArticles()
      return {
        articles: result.data,
      }
    } catch (e) {
      return {
        articles: [],
      }
    }
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.home {
  max-width: 640px;
  margin: 0 auto;
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

.article-list {
  padding: 0 15px;
}
</style>
