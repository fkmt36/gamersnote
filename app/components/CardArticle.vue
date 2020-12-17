<template>
  <NuxtLink :to="to" class="card-article">
    <div class="thumbnail">
      <img :src="thumbnail" />
    </div>
    <div class="content">
      <div class="header">
        {{ article.title }}
      </div>
      <div class="footer">
        <div class="avatar">
          <BaseAvatar :size="20" :src="avatar" />
        </div>
        <div class="username">{{ username }}</div>
        <div class="date">{{ date }}</div>
      </div>
    </div>
  </NuxtLink>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'
import { Article } from '@/api-client-axios'
import BaseAvatar from '@/components/BaseAvatar.vue'

export default Vue.extend({
  name: 'CardArticle',

  components: {
    BaseAvatar,
  },

  props: {
    article: {
      type: Object,
      required: true,
    } as PropOptions<Article>,
  },

  computed: {
    thumbnail(): string {
      return this.article.thumbnail_url
    },
    title(): string {
      return this.article.title
    },
    username(): string {
      if (this.article.author) {
        return this.article.author.username
      } else {
        return ''
      }
    },
    avatar(): string {
      if (this.article.author) {
        return this.article.author.avatar_url || ''
      } else {
        return ''
      }
    },
    date(): string {
      return this.article.created_at || ''
    },
    to(): string {
      if (this.article.article_id) {
        return '/articles/' + this.article.article_id
      } else {
        return '/'
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.card-article {
  display: flex;
  max-width: 475px;
  margin: 0 auto;
  padding: 15px 0;
  border-bottom: 1px solid $base-grey;

  .thumbnail img {
    background-color: $base-grey;
    width: 80px;
    height: 80px;
    border-radius: 10px;
    object-fit: cover;
  }

  .content {
    margin-left: 10px;
    width: 100%;
    display: flex;
    flex-direction: column;
  }

  .content .header {
    font-weight: bold;
  }

  .content .footer {
    margin: auto 0 0 0;
    font-size: $fs-tiny;

    .avatar,
    .username,
    .date {
      display: flex;
      align-items: center;
    }

    .avatar {
      margin-right: 5px;
    }

    .date {
      margin-left: auto;
    }
  }
}
.footer {
  display: flex;
}
</style>
