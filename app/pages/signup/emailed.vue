<template>
  <div class="emailed">
    <BaseHeadline text="新規登録" />
    <div class="email">
      <h2>
        <font-awesome-icon :icon="['fas', 'envelope']" />
        仮登録メールを送信しました
      </h2>
      <div class="main-message">
        <p>{{ message }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseHeadline from '@/components/BaseHeadline.vue'
import { signupState } from '@/store'

export default Vue.extend({
  components: {
    BaseHeadline,
  },

  computed: {
    message(): string {
      return `「${signupState.getEmail}」宛にno-reply@gamersnote.comからメールが届きますので、30分以内に本登録を完了してください。`
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

  .main-message {
    margin: 15px 0;
  }
}
</style>
