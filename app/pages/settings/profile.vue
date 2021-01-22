<template>
  <div>
    <BaseHeadline text="プロフィールを編集" />
    <div class="body">
      <div class="input">
        <h2>メールアドレス</h2>
        <BaseInput type="email" :on-input="onInputEmail" :text="email" />
        <p class="caption">
          メールアドレスを変更した場合、本人確認のためメールを送信します。
          30分以内に受け取ったメールアドレスにあるURLから変更を完了してください。
        </p>
      </div>
      <div class="input">
        <h2>ユーザー名 <span class="caption">※4~12字の半角英数字</span></h2>
        <BaseInput type="text" :on-input="onInputUsername" :text="username" />
        <p class="caption">
          ユーザー名を変更すると、マイページのURLが変わるので注意してください。
        </p>
      </div>
      <div class="input">
        <h2>アバター <span class="caption">※jpg、png形式（1MB以内）</span></h2>
        <div style="display: flex; align-items: center; margin-top: 5px">
          <BaseAvatar :editable="false" :src="avatarUrl" :size="60" />
          <input
            type="file"
            accept="image/png, image/jpeg"
            style="
              font-size: 0.8rem;
              margin-left: 10px;
              text-overflow: ellipsis;
              overflow: hidden;
            "
            @input="onInputImage($event)"
          />
        </div>
      </div>
      <div class="input">
        <h2>メッセージ <span class="caption">※300文字以内</span></h2>
        <div style="display: flex; align-items: center">
          <BaseTextarea :text="message" :on-input="onInputMessage" />
        </div>
      </div>
      <BaseButton text="保存" :on-click="save" :disabled="processing" />
      <div class="other-settings">
        <h2>その他の設定</h2>
        <ButtonLink to="/settings/password" text="パスワード再設定" />
        <ButtonLink to="/settings/leave" text="退会" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { meStore, baseModalState } from '@/store'
import { $userApi, $imageApi } from '@/plugins/api'
import BaseHeadline from '@/components/BaseHeadline.vue'
import BaseInput from '@/components/BaseInput.vue'
import BaseAvatar from '@/components/BaseAvatar.vue'
import BaseTextarea from '@/components/BaseTextarea.vue'
import BaseButton from '@/components/BaseButton.vue'
import ButtonLink from '@/components/ButtonLink.vue'

interface Data {
  avatarUrl: string
  username: string
  message: string
  email: string
  processing: boolean
}

export default Vue.extend({
  components: {
    BaseHeadline,
    BaseInput,
    BaseAvatar,
    BaseTextarea,
    BaseButton,
    ButtonLink,
  },

  asyncData({ redirect }): Data | void {
    const me = meStore.getMe
    if (!me) {
      return redirect('/')
    }
    return {
      avatarUrl: me.avatar_url || '/default.png',
      username: me.username,
      message: me.message || '',
      email: me.email,
      processing: false,
    }
  },

  data(): Data {
    return {} as Data
  },

  methods: {
    onInputEmail(email: string): void {
      this.email = email
    },
    onInputUsername(username: string): void {
      this.username = username
    },
    onInputMessage(message: string): void {
      this.message = message
    },
    async onInputImage(event: Event) {
      try {
        const files = (event.target as HTMLInputElement).files
        if (files && files.length) {
          const file = files[0]
          if (file.size > 1000000) {
            throw new Error('サイズオーバー')
          }
          const result = await $imageApi().uploadImage(file)
          this.avatarUrl = result.data.url
        }
      } catch (err) {
        event.preventDefault()
        baseModalState.setModal({
          showModal: true,
          message: 'エラーが発生しました',
        })
      }
    },
    updateAvatarUrl(url: string) {
      this.avatarUrl = url
    },

    async save() {
      try {
        this.processing = true
        const { data } = await $userApi().putUser({
          email: this.email,
          avatar_url: this.avatarUrl,
          username: this.username,
          message: this.message,
        })
        meStore.setMe(data)
        const message = 'プロフィールを更新しました'
        baseModalState.setModal({ showModal: true, message })
      } catch (err) {
        const message = 'プロフィールの更新に失敗しました'
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

  a {
    margin: 10px 0;
  }
}
</style>
