<template>
  <h1>{{ title }}</h1>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'PageTitle',

  props: {
    title: {
      type: String,
      required: true,
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

h1 {
  text-align: center;
  font-size: $fs-title;
  font-weight: bold;
  color: white;
  margin: 20px 0;
}
</style>
