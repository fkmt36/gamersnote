<template>
  <div>
    <CardArticle
      v-for="article in articles"
      :key="article.articleId"
      :article="article"
    />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Article } from '@/api-client-axios'
import { articlesStore } from '@/store'
import CardArticle from './CardArticle.vue'

export default Vue.extend({
  name: 'ListArticle',

  components: {
    CardArticle,
  },

  // asyncData() {
  //   console.log('hit')
  //   articlesStore.fetchAllArticles()
  // },

  computed: {
    articles(): Array<Article> {
      return articlesStore.getAllArticles
    },
  },

  mounted() {
    articlesStore.fetchAllArticles()
  },
})
</script>

<style lang="scss" scoped></style>
