<template>
  <div id="emailed">
    <PageTitle title="メールリンクを送信しました" />
    <p>
      <span>{{ email }}</span
      >宛にメールリンクを送信しました。確認をお願いいたします。
    </p>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import PageTitle from '@/components/PageTitle.vue'

export default Vue.extend({
  components: {
    PageTitle,
  },

  data() {
    return {
      email: '',
    }
  },

  mounted() {
    this.email = window.localStorage.getItem('emailForSignIn') || ''
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

#emailed {
  padding: 0 15px;

  > p:nth-last-of-type(1) {
    max-width: 310px;
    margin: 0 auto;
    color: white;
    span {
      font-weight: bold;
    }
  }
}
</style>
