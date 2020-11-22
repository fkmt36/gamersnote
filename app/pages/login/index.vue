<template>
  <div id="login">
    <PageTitle title="ログイン・新規登録" />
    <p>
      <span>メールアドレス</span>
      <Input v-model="email" :type="InputType.Email" />
    </p>
    <button @click="send">メールリンクを送信</button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import PageTitle from '@/components/PageTitle.vue'
import Input, { InputType } from '@/components/BaseInput.vue'
import { filterLoadingState } from '@/store'

export default Vue.extend({
  components: {
    PageTitle,
    Input,
  },

  data() {
    return {
      InputType,
      email: 'happyfukumoto@gmail.com',
    }
  },

  methods: {
    async send(): Promise<void> {
      try {
        filterLoadingState.setLoading({ isLoading: true, message: '処理中...' })
        await this.$fire.auth.sendSignInLinkToEmail(this.email, {
          url: 'http://localhost:3000/login/callback',
          handleCodeInApp: true,
        })
        window.localStorage.setItem('emailForSignIn', this.email)
        this.$router.push('/login/emailed')
      } catch (err) {
        console.error(err)
      } finally {
        filterLoadingState.clearLoading()
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

#login {
  padding: 0 15px;

  > p:nth-last-of-type(1) {
    max-width: 310px;
    margin: 0 auto;
    span {
      color: white;
      font-weight: bold;
    }
  }

  > button {
    display: block;
    margin: 20px auto;
    background-color: $btn-yellow;
    color: white;
    font-weight: bold;
    width: 100%;
    max-width: 250px;
    height: 30px;
    border-radius: 15px;
    text-align: center;
    box-shadow: $normal-shadow;
  }
}
</style>
