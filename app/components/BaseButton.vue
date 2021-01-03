<template>
  <button
    class="base-button"
    :disabled="disabled"
    type="button"
    @click="onClick"
  >
    {{ text }}
  </button>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'BaseAvatar',

  props: {
    text: {
      type: String,
      required: true,
    } as PropOptions<string>,
    onClick: {
      type: Function,
      required: false,
      default: () => {},
    } as PropOptions<() => void>,
    disabled: {
      type: Boolean,
      required: false,
      default: false,
    } as PropOptions<boolean>,
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.base-button {
  display: block;
  height: 40px;
  width: 100%;
  color: white;
  background-color: $btn-blue;
  border-radius: 5px;
  font-weight: bold;
  text-align: center;
}

.base-button:disabled {
  display: block;
  height: 40px;
  width: 100%;
  color: white;
  background-color: $stroke-grey;
  border-radius: 5px;
  font-weight: bold;
  text-align: center;
}
</style>
