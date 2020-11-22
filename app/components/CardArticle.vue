<template>
  <div>{{ article.title }}</div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'
import { Article } from '@/api-client'

export default Vue.extend({
  name: 'CardArticle',

  props: {
    article: {
      type: Object,
      required: true,
    } as PropOptions<Article>,
  },
})
</script>

<style lang="scss" scoped></style>
