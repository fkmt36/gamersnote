<template>
  <div class="userpage">
    <div class="header">
      <BaseAvatar :size="70" :src="avatarURL" />
      <div class="username">
        {{ username }}
      </div>
    </div>
    <div class="message">
      <p>{{ message }}</p>
    </div>
    <div v-if="isMe" class="profile-edit-btn">
      <ButtonLink text="プロフィールを編集する" to="/settings/profile" />
    </div>
    <div class="action">
      <span class="sort-title">記事</span>
      <font-awesome-icon :icon="['fas', 'sort']" />
    </div>
    <div class="article-list">
      <ListArticle :articles="articles" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $articleApi, $userApi } from '@/plugins/api'
import { meStore } from '@/store'
import BaseAvatar from '@/components/BaseAvatar.vue'
import ButtonLink from '@/components/ButtonLink.vue'
import ListArticle from '@/components/ListArticle.vue'
import { Article } from '@/api-client-axios'

interface Data {
  username: string
  avatarURL: string
  message: string
  articles: Article[]
  isMe: boolean
}

export default Vue.extend({
  components: {
    BaseAvatar,
    ButtonLink,
    ListArticle,
  },
  async asyncData({ params }): Promise<Data | void> {
    const me = meStore.getMe
    if (me && me.username === params.username) {
      try {
        const result = await $articleApi().getTheUsersArticles(params.username)
        return {
          username: me.username,
          avatarURL: me.avatar_url || '/default.png',
          message: me.message || '',
          articles: result.data,
          isMe: true,
        }
      } catch (err) {
        console.error(err)
      }
    } else {
      try {
        const result = await $userApi().getUser(params.username)
        const result2 = await $articleApi().getTheUsersArticles(params.username)
        console.log(result2.data)
        return {
          username: result.data.username,
          avatarURL: result.data.avatar_url || '/default.png',
          message: result.data.message || '',
          articles: result2.data,
          isMe: false,
        }
      } catch (err) {
        console.error(err)
      }
    }
  },
  data(): Data {
    return {} as Data
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.userpage {
  padding: 30px 20px;
  max-width: 500px;
  margin: 0 auto;
}

.header {
  display: flex;

  .username {
    margin-left: 20px;
    display: flex;
    align-items: center;
    font-size: $fs-middle;
    font-weight: bold;
  }
}

.message {
  margin: 30px 0;
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

.profile-edit-btn {
  margin: 20px 0;
}
</style>
