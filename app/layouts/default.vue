<template>
  <div>
    <TheHeader style="z-index: 1" />
    <Nuxt style="z-index: 1" />
    <FilterDarken style="z-index: 2" />
    <FilterLoading style="z-index: 3" />
    <TheMenu style="z-index: 4" />
    <BaseModal style="z-index: 4" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import TheHeader from '@/components/TheHeader.vue'
import TheMenu from '@/components/TheMenu.vue'
import FilterDarken from '@/components/FilterDarken.vue'
import FilterLoading from '@/components/FilterLoading.vue'
import BaseModal from '@/components/BaseModal.vue'

export default Vue.extend({
  components: {
    TheHeader,
    TheMenu,
    FilterDarken,
    FilterLoading,
    BaseModal,
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
  color: $font-black;
}

body {
  background-color: white;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}
</style>
