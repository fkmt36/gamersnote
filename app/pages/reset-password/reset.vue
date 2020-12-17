<template>
  <div class="signin">
    <BaseHeadline text="パスワード再設定" />
    <div class="email">
      <div class="input">
        <h3>新しいパスワード</h3>
        <BaseInput type="password" :on-input="onInputPassword" />
      </div>
      <BaseButton text="登録" :on-click="() => {}" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseButton from '@/components/BaseButton.vue'

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseButton,
  },

  data() {
    return {
      password: '',
    }
  },

  methods: {
    onInputPassword(password: string) {
      this.password = password
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
</style>
