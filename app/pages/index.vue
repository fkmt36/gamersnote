<template>
  <div>
    <div id="header-section">
      <h1>
        <div>ゲームの楽しさを</div>
        <div>文章で共有しよう</div>
      </h1>
      <div>
        <p>今すぐはじめる</p>
        <div id="login-btn">
          <nuxt-link to="/login">ログイン・新規登録</nuxt-link>
        </div>
      </div>
    </div>
    <div id="list-article-section">
      <ListArticle />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import ListArticle from '@/components/ListArticle.vue'

export default Vue.extend({
  components: {
    ListArticle,
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

#header-section {
  padding: 0 15px;

  h1 {
    font-size: $fs-title;
    font-weight: bold;
    color: white;
    margin: 20px auto;
    max-width: 310px;

    div:nth-of-type(1) {
      text-align: left;
      margin-bottom: 20px;
    }

    div:nth-of-type(2) {
      text-align: right;
    }
  }

  > div:nth-of-type(1) {
    max-width: 310px;
    margin: 50px auto;

    p {
      color: white;
      font-weight: bold;
    }

    #login-btn {
      color: white;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: $btn-yellow;
      margin: 10px auto 0 auto;
      border-radius: 20px;
      width: 100%;
      height: 35px;
      max-width: 350px;
      font-weight: bold;
      padding: 5px;
      box-shadow: $normal-shadow;
    }
  }
}

#list-article-section {
  background-color: white;
  border-radius: 10px 10px 0 0;
}
</style>
