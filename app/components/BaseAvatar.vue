<template>
  <div>
    <div v-if="editable">
      <button @click="selectImage">
        <img :src="src" :style="style" />
      </button>
      <input
        v-show="false"
        ref="inputAvatar"
        type="file"
        accept="image/png, image/jpeg"
        @change="() => {}"
      />
    </div>
    <div v-if="!editable">
      <img :src="src" :style="style" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'BaseAvatar',

  props: {
    src: {
      type: String,
      required: true,
    } as PropOptions<string>,
    size: {
      type: Number,
      required: true,
    } as PropOptions<number>,
    editable: {
      type: Boolean,
      required: false,
      default: false,
    } as PropOptions<boolean>,
    onChange: {
      type: Function,
      required: false,
      default: () => {},
    } as PropOptions<(src: string) => void>,
  },

  computed: {
    style(): string {
      return `
          width: ${this.size}px;
          height: ${this.size}px;
      `
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

img {
  object-fit: cover;
  border-radius: 5px;
  background-color: $white-blue;
}
</style>
