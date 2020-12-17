<template>
  <NuxtLink class="button-link" :to="to">
    {{ text }}
  </NuxtLink>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'ButtonLink',

  props: {
    to: {
      type: String,
      required: true,
    } as PropOptions<string>,
    text: {
      type: String,
      required: true,
    } as PropOptions<string>,
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.button-link {
  display: flex;
  height: 40px;
  width: 100%;
  color: $dark-grey;
  border: 1px solid $stroke-grey;
  border-radius: 5px;
  font-weight: bold;
  justify-content: center;
  align-items: center;
}
</style>
