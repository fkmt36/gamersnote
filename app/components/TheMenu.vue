<template>
  <transition name="right">
    <div v-show="show" class="menu-wrapper" @click.self="close">
      <nav>
        <div class="menu-header">
          <BaseAvatar
            v-show="isLogin"
            :editable="false"
            :src="avatar"
            :size="30"
          />
          <div v-show="isLogin" class="username">{{ username }}</div>
          <button @click="close">
            <font-awesome-icon icon="times" />
          </button>
        </div>
        <div class="menu-container">
          <ul v-show="!isLogin">
            <li @click="close">
              <NuxtLink to="/">ホーム</NuxtLink>
            </li>
            <li @click="close">
              <NuxtLink to="/login">ログイン・新規登録</NuxtLink>
            </li>
          </ul>
          <ul v-show="isLogin">
            <li @click="close">
              <NuxtLink to="/">ホーム</NuxtLink>
            </li>
            <li @click="close">
              <NuxtLink to="/articles/new">記事を書く</NuxtLink>
            </li>
            <li @click="close">
              <NuxtLink :to="myPageURL">マイページ</NuxtLink>
            </li>
            <!-- <li @click="close">
              <NuxtLink to="/settings/account">アカウント設定</NuxtLink>
            </li> -->
            <li @click="logout">
              <NuxtLink to="/">ログアウト</NuxtLink>
            </li>
          </ul>
        </div>
      </nav>
    </div>
  </transition>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseAvatar from '@/components/BaseAvatar.vue'
import { $userApi } from '@/plugins/api'
import { theMenuState, meStore, filterDarkenState } from '../store'

export default Vue.extend({
  name: 'TheMenu',

  components: {
    BaseAvatar,
  },

  computed: {
    show(): boolean {
      return theMenuState.getShowMenu
    },
    isLogin(): boolean {
      return !!meStore.getMe
    },
    myPageURL(): string {
      const me = meStore.getMe
      if (me) {
        return `/${me.username}`
      }
      return '/'
    },
    username(): string {
      const me = meStore.getMe
      if (me) {
        return me.username
      }
      return ''
    },
    avatar(): string {
      if (meStore.getMe && meStore.getMe.avatar_url) {
        return meStore.getMe.avatar_url
      }
      return '/default.png'
    },
  },

  methods: {
    close(): void {
      filterDarkenState.setDarken(false)
      theMenuState.setShowMenu(false)
    },
    async logout() {
      try {
        await $userApi().patchUserSignouted()
        meStore.setMe(null)
      } catch {
      } finally {
        this.close()
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';
.menu-wrapper {
  position: absolute;
  height: 100vh;
  width: 80%;
  top: 0;
  right: 0;
}
nav {
  background-color: white;
  height: 100%;
  width: 100%;
  padding: 0 15px;
}
.menu-header {
  display: flex;
  height: 45px;
  align-items: center;

  .avatar {
    width: 30px;
    height: 30px;
    background-color: $base-grey;
    border-radius: 50px;
    margin: auto 0;
  }

  .username {
    margin: auto 0 auto 10px;
    font-weight: bold;
  }

  button {
    margin: auto 0 auto auto;

    svg {
      font-size: 24px;
    }
  }
}
.menu-container {
  padding: 0 20px;

  li {
    font-weight: bold;
    text-align: center;
    margin: 30px 0;
  }
}
.right-enter-active,
.right-leave-active {
  transform: translate(0px, 0px);
  transition: transform 225ms cubic-bezier(0, 0, 0.2, 1) 0ms;
}
.right-enter,
.right-leave-to {
  transform: translateX(100vw) translateX(0px);
}
</style>
