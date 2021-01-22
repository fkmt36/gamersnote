<template>
  <div class="base-input">
    <input
      ref="input"
      :type="type"
      :value="text"
      @input="onInput($event.target.value)"
    />
    <button
      v-if="type === 'password'"
      type="button"
      class="eye"
      @click="switchPasswordVisible"
    >
      <font-awesome-icon v-show="passwordVisible" :icon="['fas', 'eye']" />
      <font-awesome-icon
        v-show="!passwordVisible"
        :icon="['fas', 'eye-slash']"
      />
    </button>
  </div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'BaseInput',

  props: {
    type: {
      type: String,
      required: false,
      default: 'text',
    } as PropOptions<string>,
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
      passwordVisible: false,
    }
  },

  methods: {
    switchPasswordVisible() {
      this.passwordVisible = !this.passwordVisible
      const e = this.$refs.input as HTMLInputElement
      if (this.passwordVisible) {
        e.type = 'text'
      } else {
        e.type = 'password'
      }
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.base-input {
  position: relative;
}

input {
  display: block;
  background-color: $white-blue;
  border-radius: 5px;
  border: 1px solid $stroke-grey;
  height: 40px;
  width: 100%;
  padding: 2px 5px;
  color: $font-black;
}

.eye {
  position: absolute;
  top: 8px;
  right: 10px;
  color: $dark-grey;
  font-size: 20px;
}
</style>
