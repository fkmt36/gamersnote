<template>
  <div>
    <BaseHeadline text="退会申請" />
    <div class="body">
      <p>
        退会処理を行うと、あなたの投稿した記事やコメントが全て削除されます。二度と元には戻せません。
      </p>
      <div class="check">
        <input
          ref="check"
          v-model="checked"
          type="checkbox"
        />上記の注意事項を確認しました
      </div>
      <BaseButton
        text="退会"
        :on-click="leave"
        :disabled="!checked || processing"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { baseModalState } from '@/store'
import { $userApi } from '@/plugins/api'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseButton from '@/components/BaseButton.vue'

interface Data {
  processing: boolean
  checked: boolean
}

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseButton,
  },

  data(): Data {
    return {
      processing: false,
      checked: false,
    }
  },

  methods: {
    async leave() {
      try {
        this.processing = true
        await $userApi().deleteUser()
        const message = '退会が完了しました'
        baseModalState.setModal({ showModal: true, message })
      } catch (err) {
        const message = '退会に失敗しました'
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

  p {
    margin: 15px 0;
    color: red;
  }

  .check {
    text-align: center;
    font-size: 0.9rem;
    margin: 30px 0;
    color: $dark-grey;

    input {
      margin-right: 10px;
    }
  }
}
</style>
