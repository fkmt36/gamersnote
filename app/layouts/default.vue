<template>
  <div>
    <TheHeader />
    <Nuxt />
    <FilterDarken />
    <FilterLoading />
    <TheMenu />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import TheHeader from '@/components/TheHeader.vue'
import TheMenu from '@/components/TheMenu.vue'
import FilterDarken from '@/components/FilterDarken.vue'
import FilterLoading from '@/components/FilterLoading.vue'

export default Vue.extend({
  components: {
    TheHeader,
    TheMenu,
    FilterDarken,
    FilterLoading,
  },
})
</script>

<style lang="scss">
@import 'assets/global';

html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: $fs-normal;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

body {
  background-color: $bg-blue;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}
</style>
