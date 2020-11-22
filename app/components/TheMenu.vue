<template>
  <transition name="right">
    <div v-show="show" class="menu-wrapper" @click.self="close">
      <nav>
        <div class="menu-header">
          <button @click="close">
            <font-awesome-icon icon="times" />
          </button>
        </div>
        <div class="menu-container">
          <ul>
            <li @click="close"><NuxtLink to="/">ホーム</NuxtLink></li>
            <li @click="close">
              <NuxtLink to="/login">ログイン・新規登録</NuxtLink>
            </li>
          </ul>
        </div>
      </nav>
    </div>
  </transition>
</template>

<script lang="ts">
import Vue from 'vue'
import { theMenuState } from '../store'

export default Vue.extend({
  name: 'TheMenu',

  computed: {
    show(): boolean {
      return theMenuState.getShowMenu
    },
  },

  methods: {
    close(): void {
      theMenuState.setShowMenu(false)
    },
  },
})
</script>

<style lang="scss" scoped>
.menu-wrapper {
  position: absolute;
  height: 100vh;
  width: 100vw;
  top: 0;
  right: 0;
}
nav {
  background-color: white;
  height: 100vh;
  width: 80vw;
  position: absolute;
  top: 0;
  right: 0;
  padding: 0 15px;
}
.menu-header {
  display: flex;
  height: 45px;

  button {
    margin: auto 0 auto auto;

    svg {
      font-size: 24px;
    }
  }
}
.menu-container {
  padding: 20px;

  li {
    font-weight: bold;
    text-align: center;
    margin-bottom: 20px;
  }
}
.right-enter-active,
.right-leave-active {
  transform: translate(0px, 0px);
  transition: transform 225ms cubic-bezier(0, 0, 0.2, 1) 0ms;
}
.right-enter,
.right-leave-to {
  transform: translateX(100vw) translateX(0px);
}
</style>
