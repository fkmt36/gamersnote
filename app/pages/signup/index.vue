<template>
  <div class="signin">
    <BaseHeadline text="新規登録" />
    <div class="email">
      <h2>
        <font-awesome-icon :icon="['fas', 'envelope']" />
        メールアドレスで登録
      </h2>
      <form>
        <div class="input">
          <h3>ユーザー名 <span class="caption">※4~12字の半角英数字</span></h3>
          <BaseInput type="text" :on-input="onInputUsername" />
          <p v-show="invalidUsername" class="warning">
            4~12字の半角英数字である必要があります
          </p>
          <p v-show="validUsername" class="ok">OK</p>
        </div>
        <div class="input">
          <h3>メールアドレス</h3>
          <BaseInput type="email" :on-input="onInputEmail" />
          <p v-show="invalidEmail" class="warning">
            正しいメールアドレスを入力してください
          </p>
          <p v-show="validEmail" class="ok">OK</p>
        </div>
        <div class="input">
          <h3>パスワード <span class="caption">※6~20字の半角英数字</span></h3>
          <BaseInput type="password" :on-input="onInputPassword" />
          <p v-show="invalidPassword" class="warning">
            6~20字の半角英数字である必要があります
          </p>
          <p v-show="validPassword" class="ok">OK</p>
        </div>
        <BaseButton
          text="登録"
          :on-click="signup"
          :disabled="!submitable || processing"
        />
      </form>
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
import { baseModalState, signupState, meStore } from '@/store'
import { isAxiosError } from '@/utils/is-axios-error'

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseButton,
  },

  fetch({ redirect }) {
    const me = meStore.getMe
    if (me !== null) {
      return redirect('/')
    }
  },

  data() {
    return {
      username: '',
      email: '',
      password: '',
      processing: false,
    }
  },

  computed: {
    invalidUsername(): boolean {
      return !!this.username && !this.username.match(/[A-Za-z0-9\\_]{4,12}/)
    },
    invalidPassword(): boolean {
      return !!this.password && !this.password.match(/[A-Za-z0-9\\_]{6,20}/)
    },
    invalidEmail(): boolean {
      return (
        !!this.email &&
        !this.email.match(/[A-Za-z0-9-._]+@[A-Za-z0-9-._]+.[A-Za-z]+/)
      )
    },
    validUsername(): boolean {
      return !this.invalidUsername && !!this.username
    },
    validPassword(): boolean {
      return !this.invalidPassword && !!this.password
    },
    validEmail(): boolean {
      return !this.invalidEmail && !!this.email
    },
    submitable(): boolean {
      return this.validUsername && this.validPassword && this.validEmail
    },
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
        this.processing = true
        await $userApi().postUser({
          username: this.username,
          email: this.email,
          password: this.password,
        })
        signupState.setEmail(this.email)
        this.$router.push('/signup/emailed')
      } catch (err) {
        if (isAxiosError(err) && err.response) {
          const message = (err.response.data as any).message as string
          baseModalState.setModal({ showModal: true, message })
        }
      } finally {
        this.processing = false
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

    .warning {
      color: red;
    }

    .ok {
      color: green;
    }
  }
}

.footer {
  font-size: $fs-tiny;
  text-align: center;
}
</style>
