<template>
  <div></div>
</template>

<script lang="ts">
import Vue from 'vue'
import { filterLoadingState } from '@/store'

export default Vue.extend({
  data() {
    return {
      showInputEmail: false,
      inputEmail: '',
    }
  },

  middleware({ route, redirect, $fire }) {
    if (!$fire.auth.isSignInWithEmailLink(route.fullPath)) {
      redirect('/')
    }
  },

  created() {
    filterLoadingState.setLoading({ isLoading: true, message: '処理中...' })
  },

  async mounted() {
    try {
      const email =
        window.localStorage.getItem('emailForSignIn') ||
        window.prompt('確認のためにメールアドレスを入力してください') ||
        ''
      const user = await this.$fire.auth.signInWithEmailLink(
        email,
        window.location.href
      )
      window.localStorage.removeItem('emailForSignIn')
      if (
        user &&
        user.additionalUserInfo &&
        user.additionalUserInfo.isNewUser
      ) {
        // TODO: user作成APIを叩く
        // TODO: 帰ってくるuserをstoreに保存
      }
    } catch (err) {
      console.error(err)
    } finally {
      this.$router.push('/')
      filterLoadingState.clearLoading()
    }
  },
})
</script>

<style lang="scss" scoped></style>
