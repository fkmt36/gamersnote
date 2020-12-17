<template>
  <div class="signin">
    <BaseHeadline text="新規登録" />
    <div class="email">
      <h2>
        <font-awesome-icon :icon="['fas', 'envelope']" />
        メールアドレスで登録
      </h2>
      <div class="input">
        <h3>ユーザー名 <span class="caption">※4~12字の半角英数字</span></h3>
        <BaseInput type="text" :on-input="onInputUsername" />
      </div>
      <div class="input">
        <h3>メールアドレス</h3>
        <BaseInput type="email" :on-input="onInputEmail" />
      </div>
      <div class="input">
        <h3>パスワード <span class="caption">※6~20字の半角英数字</span></h3>
        <BaseInput type="password" :on-input="onInputPassword" />
      </div>
      <BaseButton text="登録" :on-click="signup" />
    </div>
    <div class="footer">
      <NuxtLink to="/signin">アカウントをお持ちの方</NuxtLink>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseButton from '@/components/BaseButton.vue'
import { $userApi } from '@/plugins/api'
import { meStore } from '@/store'

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseButton,
  },

  data() {
    return {
      username: '',
      email: '',
      password: '',
    }
  },

  methods: {
    onInputUsername(username: string) {
      this.username = username
    },
    onInputEmail(email: string) {
      this.email = email
    },
    onInputPassword(password: string) {
      this.password = password
    },
    async signup() {
      try {
        const result = await $userApi().postUser({
          username: this.username,
          email: this.email,
          password: this.password,
        })
        meStore.setMe(result.data)
        this.$router.push('/')
      } catch (err) {
        console.error(err)
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.email {
  padding: 15px 10px;
  max-width: 450px;
  margin: 0 auto;

  h2 {
    color: $dark-grey;
    font-weight: bold;
  }

  .input {
    color: $dark-grey;
    margin: 30px 0;

    .caption {
      font-size: $fs-tiny;
    }
  }
}

.footer {
  font-size: $fs-tiny;
  text-align: center;
}
</style>
