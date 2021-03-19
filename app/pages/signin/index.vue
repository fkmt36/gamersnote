<template>
  <div class="signin">
    <BaseHeadline text="ログイン" />
    <div class="email">
      <h2>
        <font-awesome-icon :icon="['fas', 'envelope']" />
        メールアドレスでログイン
      </h2>
      <form>
        <div class="input">
          <h3>メールアドレス</h3>
          <BaseInput type="email" :on-input="onInputEmail" />
        </div>
        <div class="input">
          <h3>パスワード</h3>
          <BaseInput type="password" :on-input="onInputPassword" />
        </div>
      </form>
      <BaseButton
        text="ログイン"
        :disabled="!submitable || processing"
        :on-click="signin"
      />
    </div>
    <div class="footer">
      <NuxtLink to="/reset-password/request">パスワードを忘れた方</NuxtLink>
      <span class="partition">|</span>
      <NuxtLink to="/signup">アカウントをお持ちでない方</NuxtLink>
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
import { baseModalState } from '~/store'

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
      email: '',
      password: '',
      processing: false,
    }
  },

  computed: {
    submitable(): boolean {
      return !!this.email && !!this.password
    },
  },

  methods: {
    onInputEmail(email: string) {
      this.email = email
    },
    onInputPassword(password: string) {
      this.password = password
    },
    async signin() {
      try {
        this.processing = true
        const result = await $userApi().patchUserSignined({
          email: this.email,
          password: this.password,
        })
        meStore.setMe(result.data)
        const to = this.$route.query.from as string
        if (to === undefined || to === null || to === '') {
          await this.$router.push('/')
        } else {
          await this.$router.push(to)
        }
      } catch (err) {
        const message = 'ログインに失敗しました'
        baseModalState.setModal({ showModal: true, message })
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
  }
}

.footer {
  font-size: $fs-tiny;
  text-align: center;
}
</style>
