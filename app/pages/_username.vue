<template>
  <div class="userpage">
    <div class="header">
      <BaseAvatar :size="70" :src="avatarURL" />
      <div class="username">
        {{ username }}
      </div>
    </div>
    <div class="message">
      <p>{{ message }}</p>
    </div>
    <div class="profile-edit-btn">
      <ButtonLink text="プロフィールを編集する" to="/settings/profile" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $userApi } from '@/plugins/api'
import { meStore } from '@/store'
import BaseAvatar from '@/components/BaseAvatar.vue'
import ButtonLink from '@/components/ButtonLink.vue'

interface Data {
  username: string
  avatarURL: string
  message: string
}

export default Vue.extend({
  components: {
    BaseAvatar,
    ButtonLink,
  },
  async asyncData({ params }): Promise<Data | void> {
    const me = meStore.getMe
    if (me && me.username === params.username) {
      return {
        username: me.username,
        avatarURL: me.avatar_url || '',
        message: me.message || '',
      }
    }
    try {
      const result = await $userApi().getUser(params.username)
      return {
        username: result.data.username,
        avatarURL: result.data.avatar_url || '',
        message: result.data.message || '',
      }
    } catch (err) {}
  },
  data(): Data {
    return {} as Data
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.userpage {
  padding: 30px 20px;
  max-width: 500px;
  margin: 0 auto;
}

.header {
  display: flex;

  .username {
    margin-left: 20px;
    display: flex;
    align-items: center;
    font-size: $fs-middle;
    font-weight: bold;
  }
}

.message {
  margin: 30px 0;
}

.profile-edit-btn {
  margin: 20px 0;
}
</style>
