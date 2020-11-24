<template>
  <input
    :type="type"
    :value="value"
    @input="$emit('input', $event.target.value)"
  />
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export enum InputType {
  Email = 'email',
  Password = 'password',
  Text = 'text',
}

export default Vue.extend({
  name: 'BaseInput',

  props: {
    type: {
      type: String,
      required: true,
    } as PropOptions<InputType>,
    value: {
      type: String,
      required: true,
    },
  },
})
</script>

<style lang="scss" scoped>
// @import 'assets/global';

input {
  color: black;
  background-color: white;
  width: 100%;
  font-weight: normal;
  border-radius: 10px;
  border: 1px solid grey;
  height: 30px;
  padding: 2px;
}
</style>
