<template>
  <div></div>
</template>

<script lang="ts">
import Vue from 'vue'
import { filterLoadingState, meStore, authState } from '@/store'

export default Vue.extend({
  middleware({ route, redirect, $fire }) {
    if (!$fire.auth.isSignInWithEmailLink(route.fullPath)) {
      redirect('/')
    }
  },

  created() {
    filterLoadingState.setLoading({ isLoading: true, message: '処理中...' })
  },

  async mounted() {
    await authState.verifyEmailFromEmailLink(
      this.getEmail(),
      window.location.href
    )
    await meStore.fetchMe()
    if (authState.getToken && meStore.getMe) {
      this.$router.push('/')
    } else if (authState.getToken && !meStore.getMe) {
      this.$router.push('/welcome')
    } else {
      // エラースタックの処理
      this.$router.push('/')
    }
    filterLoadingState.clearLoading()
  },
  methods: {
    getEmail(): string {
      // TODO: メールアドレス再入力を強化
      return (
        window.localStorage.getItem('emailForSignIn') ||
        window.prompt('確認のためにメールアドレスを入力してください') ||
        ''
      )
    },
  },
})
</script>

<style lang="scss" scoped></style>
