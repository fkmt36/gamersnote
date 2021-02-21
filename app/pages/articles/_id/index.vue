<template>
  <div>
    <article>
      <div class="thumbnail">
        <img :src="thumbnailUrl" />
      </div>
      <div class="content">
        <div class="title">
          <h1>{{ title }}</h1>
        </div>
        <div class="information">
          <BaseAvatar :src="avatarURL" :size="30" />
          <div>
            <div class="username">{{ username }}</div>
            <div class="date">{{ createdAt }}</div>
          </div>
          <div class="good">
            <font-awesome-icon :icon="['fas', 'heart']" />
            <span>{{ likeCount }}</span>
          </div>
        </div>
        <div class="article-body" v-html="body"></div>
      </div>
      <div v-if="showLikeBtn">
        <button v-show="!liked" class="like-btn" @click="like"></button>
        <button v-show="liked" class="liked-btn" @click="like"></button>
      </div>
    </article>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $articleApi, $likeApi } from '@/plugins/api'
import { meStore } from '@/store'
import BaseAvatar from '@/components/BaseAvatar.vue'

interface Data {
  articleId: string
  thumbnailUrl: string
  title: string
  body: string
  avatarURL: string
  username: string
  createdAt: string
  likeCount: number
  liked: boolean
}

export default Vue.extend({
  components: {
    BaseAvatar,
  },
  async asyncData({ params, redirect, req }): Promise<Data | void> {
    try {
      // TODO: 並列実行
      const { data } = await $articleApi().getArticle(params.id)
      if (!data) {
        redirect('/')
      }
      const liked = await (async (): Promise<boolean> => {
        try {
          if (process.server) {
            await $likeApi().getLike(params.id, {
              headers: {
                Cookie: req.headers.cookie,
              },
            })
          } else {
            await $likeApi().getLike(params.id)
          }
          return true
        } catch {
          return false
        }
      })()
      return {
        articleId: params.id,
        thumbnailUrl: data.thumbnail_url,
        title: data.title,
        body: data.body,
        avatarURL: data.author
          ? data.author.avatar_url || '/default.png'
          : '/default.png',
        username: data.author ? data.author.username : '',
        createdAt: data.created_at ? data.created_at : '',
        likeCount: data.like_count ? data.like_count : 0,
        liked,
      }
    } catch {
    } finally {
    }
  },
  data(): Data {
    return {} as Data
  },
  computed: {
    showLikeBtn(): boolean {
      return meStore.getMe !== null
    },
  },
  methods: {
    async like() {
      try {
        if (!this.liked) {
          await $likeApi().putLike(this.articleId)
          this.likeCount += 1
        } else {
          await $likeApi().deleteLike(this.articleId)
          this.likeCount -= 1
        }
        this.liked = !this.liked
      } catch {}
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

article {
  max-width: 640px;
  margin: 0 auto;

  .thumbnail {
    width: 100%;
    padding-top: 56.25%;
    overflow: hidden;
    position: relative;

    img {
      position: absolute;
      top: 50%;
      left: 50%;
      width: 100%;
      transform: translate(-50%, -50%);
    }

    @media screen and (min-width: 640px) {
      margin: 50px 0;
    }
  }

  .content {
    padding: 0 10px;
  }

  .title {
    margin: 30px 0;

    h1 {
      font-weight: bold;
      font-size: $fs-title;
    }
  }

  .information {
    display: flex;
    margin: 30px 0;

    .username,
    .date {
      font-size: 0.8rem;
      margin-left: 10px;
      color: $dark-grey;
    }

    .username {
      font-weight: bold;
    }

    .good {
      margin: auto 0 auto auto;
      color: $good-color;
    }
  }
  .like-btn,
  .liked-btn {
    width: 60px;
    height: 60px;
    background-size: contain;
    border-radius: 50%;
    position: fixed;
    right: 25px;
    bottom: 25px;
    box-shadow: 0 15px 30px -5px rgb(0 0 0 / 15%), 0 0 5px rgb(0 0 0 / 10%);
  }
  .like-btn {
    background-image: url('/like.png');
  }
  .liked-btn {
    background-image: url('/liked.png');
  }
}
</style>

<style>
.article-body p {
  margin: 10px 0;
}
.article-body h2 {
  margin: 10px 0;
  font-size: 1.3rem;
  font-weight: bold;
}
.article-body img {
  display: block;
  max-width: 80%;
  max-height: 56.25%;
  margin: 0 auto;
}
.article-body .iframe-wrapper {
  position: relative;
  padding-bottom: 56.25%;
  height: 0;
  overflow: hidden;
}
.article-body iframe {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>
