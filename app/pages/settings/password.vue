<template>
  <div>
    <BaseHeadline text="パスワード設定" />
    <div class="body">
      <div class="input">
        <h2>
          新しいパスワード <span class="caption">※6~20字の半角英数字</span>
        </h2>
        <BaseInput type="password" :on-input="onInputPassword" />
      </div>
      <BaseButton text="更新" :on-click="save" :disabled="processing" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { baseModalState, meStore } from '@/store'
import { $userApi } from '@/plugins/api'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseButton from '@/components/BaseButton.vue'

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseButton,
  },

  fetch({ redirect, route }) {
    const me = meStore.getMe
    if (me === null) {
      redirect('/signin?from=' + route.path)
    }
  },

  data() {
    return {
      password: '',
      processing: false,
    }
  },

  methods: {
    onInputPassword(password: string): void {
      this.password = password
    },
    async save() {
      try {
        this.processing = true
        await $userApi().putPassword({
          password: this.password,
        })
        const message = 'パスワードを更新しました'
        baseModalState.setModal({ showModal: true, message })
      } catch (err) {
        const message = 'パスワードの更新に失敗しました'
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

.body {
  max-width: 450px;
  margin: 0 auto;
  padding: 15px 10px;
}

.input {
  color: $dark-grey;
  margin: 30px 0;

  .caption {
    font-size: $fs-tiny;
  }
}

.other-settings {
  margin: 40px 0;

  h2 {
    color: $dark-grey;
    font-weight: bold;
  }

  button {
    background-color: white;
    color: $dark-grey;
    border: 1px solid $stroke-grey;
    margin: 10px 0;
  }
}
</style>
