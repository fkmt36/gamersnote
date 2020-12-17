<template>
  <div class="list-article">
    <CardArticle
      v-for="article in articles"
      :key="article.articleId"
      :article="article"
    />
  </div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'
import { Article } from '@/api-client-axios'
import CardArticle from './CardArticle.vue'

export default Vue.extend({
  name: 'ListArticle',

  components: {
    CardArticle,
  },

  props: {
    articles: {
      type: Array,
      required: true,
    } as PropOptions<Array<Article>>,
  },
})
</script>

<style lang="scss" scoped></style>
