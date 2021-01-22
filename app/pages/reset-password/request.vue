<template>
  <div class="signin">
    <BaseHeadline text="パスワード再設定リクエストを送る" />
    <div class="message">
      <p>
        登録中のメールアドレスを入力して「送信」ボタンをクリックしてください。
        パスワードの再設定方法についてのメールが届きますので、メールに記載の指示に従ってください。
      </p>
    </div>
    <div class="email">
      <div class="input">
        <h3>メールアドレス</h3>
        <BaseInput type="email" :on-input="onInputEmail" />
      </div>
      <BaseButton text="送信" :on-click="requestRestPassword" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseButton from '@/components/BaseButton.vue'
import { $userApi } from '@/plugins/api'
import { baseModalState } from '@/store'

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseButton,
  },

  data() {
    return {
      email: '',
      processing: false,
    }
  },

  methods: {
    onInputEmail(email: string) {
      this.email = email
    },

    async requestRestPassword() {
      try {
        this.processing = true
        await $userApi().patchUserPasswordReset({
          email: this.email,
        })
        const message = 'メールを送信しました'
        baseModalState.setModal({ showModal: true, message })
      } catch (e) {
        const message = 'メールの送信に失敗しました'
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

.message {
  max-width: 450px;
  margin: 0 auto;
  padding: 15px 10px 0 10px;
}

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
</style>
