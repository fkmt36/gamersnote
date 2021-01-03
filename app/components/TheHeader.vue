<template>
  <header>
    <div class="logo">
      <NuxtLink to="/">GamersNote</NuxtLink>
    </div>
    <div v-if="authenticated" class="authenticated-action">
      <button type="button" @click="show">
        <BaseAvatar :src="avatar" :size="30" />
      </button>
    </div>
    <div v-else class="unauthenticated-action">
      <NuxtLink to="/signin">ログイン</NuxtLink>
      <span class="partition">|</span>
      <NuxtLink to="/signup">新規登録</NuxtLink>
    </div>
  </header>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseAvatar from '@/components/BaseAvatar.vue'
import { theMenuState, filterDarkenState, meStore } from '@/store'

export default Vue.extend({
  name: 'TheHeader',

  components: {
    BaseAvatar,
  },

  computed: {
    authenticated(): boolean {
      return !!meStore.getMe
    },
    avatar(): string {
      if (meStore.getMe && meStore.getMe.avatar_url) {
        return meStore.getMe.avatar_url
      }
      return '/default.png'
    },
  },

  methods: {
    show(): void {
      filterDarkenState.setDarken(true)
      theMenuState.setShowMenu(true)
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

header {
  display: flex;
  height: 40px;
  padding: 0 10px;
  background-color: $bg-blue;
}

.logo,
.unauthenticated-action,
.authenticated-action {
  color: white;
  display: flex;
  align-items: center;
}

.logo {
  font-size: 18px;
  font-weight: bold;
}

.unauthenticated-action,
.authenticated-action {
  margin-left: auto;

  .partition {
    margin: 0 10px;
  }
}
</style>
