<template>
  <header>
    <NuxtLink to="/">GamersNote</NuxtLink>
    <button @click="switchShowMenu">
      <font-awesome-icon icon="bars" />
    </button>
  </header>
</template>

<script lang="ts">
import Vue from 'vue'
import { theMenuState } from '@/store'

export default Vue.extend({
  name: 'TheHeader',

  methods: {
    switchShowMenu(): void {
      theMenuState.setShowMenu(!theMenuState.getShowMenu)
    },
  },
})
</script>

<style lang="scss" scoped>
header {
  display: flex;
  height: 45px;
  padding: 0 10px;
}
a {
  display: block;
  margin: auto 0;
  color: white;
  text-decoration: none;
  font-size: 18px;
  font-weight: bold;
}
button {
  margin: auto 0 auto auto;
}
svg {
  font-size: 24px !important;
  color: white;
}
</style>
