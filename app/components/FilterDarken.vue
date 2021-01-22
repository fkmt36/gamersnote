<template>
  <transition name="fade">
    <div v-show="show"></div>
  </transition>
</template>

<script lang="ts">
import Vue from 'vue'
import { filterDarkenState } from '@/store'

export default Vue.extend({
  name: 'FilterDarken',

  computed: {
    show(): boolean {
      return filterDarkenState.getIsDarken
    },
  },
})
</script>

<style lang="scss" scoped>
div {
  background-color: rgba(0, 0, 0, 0.4);
  position: fixed;
  height: 100vh;
  width: 100vw;
  top: 0;
  right: 0;
}
</style>
