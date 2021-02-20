<template>
  <div class="card-article">
    <NuxtLink class="article-link" :to="to"></NuxtLink>
    <div class="thumbnail">
      <img :src="thumbnail" />
    </div>
    <div class="content">
      <div class="header">
        {{ article.title }}
      </div>
      <div class="footer">
        <NuxtLink class="user" :to="toUserPage">
          <div class="avatar">
            <BaseAvatar :size="20" :src="avatar" />
          </div>
          <div class="username">{{ username }}</div>
        </NuxtLink>
        <div class="date">{{ date }}</div>
        <div v-show="ownArticle" class="modal-wrapper">
          <button class="dotbtn" @click.stop="switchModal">
            <font-awesome-icon :icon="['fas', 'ellipsis-v']" />
          </button>
          <div v-show="openModal" class="modal">
            <div><NuxtLink :to="toArticleEdit">記事を編集する</NuxtLink></div>
            <div><button @click="openModal = false">キャンセル</button></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'
import { Article } from '@/api-client-axios'
import { meStore } from '@/store'
import BaseAvatar from '@/components/BaseAvatar.vue'

export default Vue.extend({
  name: 'CardArticle',

  components: {
    BaseAvatar,
  },

  data() {
    return {
      openModal: false,
    }
  },

  props: {
    article: {
      type: Object,
      required: true,
    } as PropOptions<Article>,
  },

  computed: {
    thumbnail(): string {
      return this.article.thumbnail_url || '/default.png'
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
      const author = this.article.author
      if (author && author.avatar_url) {
        return author.avatar_url
      }
      return '/default.png'
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
    toUserPage(): string {
      return '/' + this.username
    },
    toArticleEdit(): string {
      return `/articles/${this.article.article_id}/edit`
    },
    ownArticle(): boolean {
      if (meStore.getMe && this.article.author) {
        return this.article.author.user_id === meStore.getMe.user_id
      }
      return false
    },
  },
  methods: {
    switchModal() {
      this.openModal = !this.openModal
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
  position: relative;

  .article-link {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }

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

    .dotbtn {
      display: flex;
      margin: 0 8px;
      align-items: center;
      height: 100%;
    }
  }
}
.footer {
  display: flex;

  .user {
    display: flex;
    z-index: 2;
  }
}
.modal-wrapper {
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
</style>
