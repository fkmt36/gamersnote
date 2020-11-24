<template>
  <div class="welcome">
    <PageTitle title="ようこそGamersNoteへ！" />
    <div class="message">
      <p>プロフィールを設定しましょう</p>
    </div>
    <form class="form">
      <p>
        <span>ユーザー名</span>
        <BaseInput v-model="username" type="text" />
      </p>
      <p>
        <span>GamersNoteID</span>
        <BaseInput v-model="gamersNoteID" type="text" />
      </p>
      <button type="button" @click="sendProfile">送信</button>
    </form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { meStore } from '@/store'
import PageTitle from '@/components/PageTitle.vue'
import BaseInput from '@/components/BaseInput.vue'

interface Params {
  gamersNoteID: string
  username: string
}

export default Vue.extend({
  components: {
    PageTitle,
    BaseInput,
  },
  data(): Params {
    return {
      gamersNoteID: '',
      username: '',
    }
  },
  methods: {
    async sendProfile() {
      await meStore.postMe({
        gamersnote_id: this.gamersNoteID,
        username: this.username,
        avatar_url: '',
        message: '',
      })
      if (meStore.getMe) {
        this.$router.push('/')
      } else {
        // TODO: エラースタックを見に行き、エラーの対応
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.welcome {
  width: 90%;
  max-width: 375px;
  margin: 0 auto;

  .message {
    color: white;
    text-align: center;
    margin: 40px 0;
  }

  .form {
    p {
      margin: 15px 0;
      span {
        color: white;
      }
    }

    button {
      background-color: $btn-yellow;
      color: white;
      margin: 0 auto;
      display: block;
      width: 80px;
      height: 25px;
      border-radius: 10px;
      font-weight: bold;
      text-align: center;
      box-shadow: $normal-shadow;
    }
  }
}
</style>
