<template>
  <div>
    <BaseHeadline text="プロフィールを編集" />
    <div class="body">
      <div class="input">
        <h2>ユーザー名 <span class="caption">※4~12字の半角英数字</span></h2>
        <BaseInput type="text" :on-input="onInputUsername" />
      </div>
      <div class="textarea">
        <textarea></textarea>
      </div>
    </div>
    <!-- <form class="form">
      <div class="section">
        <div class="input-title">プロフィール画像</div>
        <div class="input-subtitle">pngまたはjpegで10MB以下</div>
        <div class="select-avatar">
          <BaseAvatar
            :size="65"
            :src="avatarUrl"
            :editable="true"
            :on-change="updateAvatarUrl"
          />
        </div>
      </div>

      <div class="section">
        <div class="input-title">ユーザー名</div>
        <div class="input-subtitle">4~20文字</div>
        <div class="input-username">
          <BaseInput v-model="username" type="text" />
        </div>
      </div>

      <div class="section">
        <div class="input-title">GamersNoteID</div>
        <div class="input-subtitle">4~12文字の半角英数字</div>
        <div class="input-gamersnoteid">
          <BaseInput v-model="gamersNoteId" type="text" />
        </div>
      </div>

      <div class="btn-section">
        <button type="button" class="cancel-btn" @click="cancel">
          キャンセル
        </button>
        <button type="button" class="save-btn" @click="save">保存</button>
      </div>
    </form> -->
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { meStore, filterLoadingState } from '@/store'
import { $userApi } from '@/plugins/api'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseAvatar from '@/components/BaseAvatar.vue'

interface Data {
  avatarUrl: string
  username: string
  message: string
}

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseAvatar,
  },

  asyncData({ redirect }): Data | void {
    const me = meStore.getMe
    if (!me) {
      return redirect('/')
    }
    return {
      avatarUrl: me.avatar_url || '',
      username: me.username,
      message: me.message || '',
    }
  },

  data(): Data {
    return {} as Data
  },

  methods: {
    updateAvatarUrl(url: string) {
      this.avatarUrl = url
    },

    cancel() {
      this.$router.push('/')
    },

    async save() {
      try {
        filterLoadingState.setLoading({ isLoading: true, message: '処理中...' })
        await $userApi().putUser({
          email: meStore.getMe ? meStore.getMe.email : '',
          avatar_url: this.avatarUrl,
          username: this.username,
          message: this.message,
        })
      } catch (err) {
      } finally {
        filterLoadingState.clearLoading()
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
}

.input {
  color: $dark-grey;
  margin: 30px 0;

  .caption {
    font-size: $fs-tiny;
  }
}
</style>
