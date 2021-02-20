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
    </article>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $articleApi } from '@/plugins/api'
import BaseAvatar from '@/components/BaseAvatar.vue'

interface Data {
  thumbnailUrl: string
  title: string
  body: string
  avatarURL: string
  username: string
  createdAt: string
  likeCount: number
}

export default Vue.extend({
  components: {
    BaseAvatar,
  },
  async asyncData({ params, redirect }): Promise<Data | void> {
    try {
      const { data } = await $articleApi().getArticle(params.id)
      if (!data) {
        redirect('/')
      }
      return {
        thumbnailUrl: data.thumbnail_url,
        title: data.title,
        body: data.body,
        avatarURL: data.author
          ? data.author.avatar_url || '/default.png'
          : '/default.png',
        username: data.author ? data.author.username : '',
        createdAt: data.created_at ? data.created_at : '',
        likeCount: data.like_count ? data.like_count : 0,
      }
    } catch {
    } finally {
    }
  },
  data(): Data {
    return {} as Data
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
