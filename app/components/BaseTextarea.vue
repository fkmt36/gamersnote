<template>
  <textarea v-model="tmp" @input="onInput($event.target.value)"></textarea>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'BaseInput',

  props: {
    onInput: {
      type: Function,
      required: false,
      default: () => {},
    } as PropOptions<(v: string) => void>,
    text: {
      type: String,
      required: false,
      default: '',
    } as PropOptions<string>,
  },

  data() {
    return {
      tmp: '',
    }
  },

  beforeMount() {
    this.tmp = this.text
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

textarea {
  display: block;
  background-color: $white-blue;
  border-radius: 5px;
  border: 1px solid $stroke-grey;
  height: 150px;
  width: 100%;
  padding: 2px 5px;
  color: $font-black;
}
</style>
