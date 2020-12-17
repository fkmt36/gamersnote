<template>
  <div class="base-headline">
    <h1>{{ text }}</h1>
  </div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'BaseHeadline',

  props: {
    text: {
      type: String,
      required: true,
    } as PropOptions<string>,
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.base-headline {
  height: 50px;
  background-color: $white-blue;
  display: flex;
  align-items: center;
  padding: 0 10px;

  h1 {
    color: $dark-grey;
    font-weight: bold;
    font-size: 18px;
  }
}
</style>
